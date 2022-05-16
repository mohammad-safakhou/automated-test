package cmd

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"test-manager/handlers"
	"test-manager/repos"
	"test-manager/repos/influx"
	"test-manager/tasks/push"
	"test-manager/utils"
	"time"
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))

		redisClient, err := utils.CreateRedisConnection(context.TODO(), "localhost", "6379", 3*time.Second)
		if err != nil {
			panic(err)
		}

		psqlDb, err := utils.PostgresConnection("localhost", "5432", "root", "root", "tester", "disable")
		if err != nil {
			panic(err)
		}
		asynqClient := asynq.NewClient(asynq.RedisClientOpt{
			Addr:        redisClient.Options().Addr,
			DialTimeout: redisClient.Options().DialTimeout,
			Username:    redisClient.Options().Username,
			Password:    redisClient.Options().Password,
		})
		taskPusher := push.NewTaskPush(asynqClient)
		influxClient, writeAPI, queryAPI, err := utils.CreateInfluxDBConnection(context.TODO(), "aiNlFChQ9RswcCapCtLZUnH2QkleksShwQvnrtTW7obAmh0W5bW7yiLqyQwrX-pSpQc0yUFliW0hgdd4kdk96A==", "http://localhost:8086", "test", "my-bucket")
		if err != nil {
			panic(err)
		}
		defer influxClient.Close()

		endpointRepo := repos.NewEndpointRepository(psqlDb)
		netCatRepo := repos.NewNetCatRepository(psqlDb)
		pageSpeedRepo := repos.NewPageSpeedRepository(psqlDb)
		pingRepo := repos.NewPingRepository(psqlDb)
		traceRouteRepo := repos.NewTraceRouteRepository(psqlDb)
		dataCenterRepo := repos.NewDataCentersRepositoryRepository(psqlDb)
		endpointReportRepo := influx.NewEndpointReportRepository(writeAPI, queryAPI, psqlDb)

		agentHandler := handlers.NewAgentHandler()
		//endpointHandler := handlers.NewEndpointHandler(endpointRepo, dataCenterRepo, taskPusher, agentHandler)
		ruleHandler := handlers.NewRulesHandler(endpointRepo, netCatRepo, pageSpeedRepo, pingRepo, traceRouteRepo, dataCenterRepo, taskPusher, agentHandler)
		controllers := handlers.NewHttpControllers(ruleHandler, endpointReportRepo)

		e.GET("/", controllers.Hello)
		e.POST("/rules/register", controllers.RegisterRules)
		e.POST("/report/endpoint/:project_id", controllers.ReportEndpoint)

		// Start server
		go func() {
			if err := e.Start(":10000"); err != nil && err != http.ErrServerClosed {
				log.Fatal("shutting down server")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	},
}
