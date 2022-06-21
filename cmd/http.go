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

		accountRepo := repos.NewAccountsRepositoryRepository(psqlDb)
		projectRepo := repos.NewProjectsRepositoryRepository(psqlDb)
		endpointRepo := repos.NewEndpointRepository(psqlDb)
		netCatRepo := repos.NewNetCatRepository(psqlDb)
		pageSpeedRepo := repos.NewPageSpeedRepository(psqlDb)
		pingRepo := repos.NewPingRepository(psqlDb)
		traceRouteRepo := repos.NewTraceRouteRepository(psqlDb)
		aggregateRepo := repos.NewAggregateRepository(psqlDb, endpointRepo, netCatRepo, pageSpeedRepo, pingRepo, traceRouteRepo)
		packageRepo := repos.NewPackagesRepository(psqlDb)
		dataCenterRepo := repos.NewDataCentersRepositoryRepository(psqlDb)
		endpointReportRepo := influx.NewEndpointReportRepository(writeAPI, queryAPI, psqlDb)
		netCatReportRepo := influx.NewNetCatsReportRepository(writeAPI, queryAPI, psqlDb)
		pageSPeedReportRepo := influx.NewPageSpeedReportRepository(writeAPI, queryAPI, psqlDb)
		pingReportRepo := influx.NewPingReportRepository(writeAPI, queryAPI, psqlDb)
		traceRouteReportRepo := influx.NewTraceRouteReportRepository(writeAPI, queryAPI, psqlDb)

		agentHandler := handlers.NewAgentHandler()
		//endpointHandler := handlers.NewEndpointHandler(endpointRepo, dataCenterRepo, taskPusher, agentHandler)
		ruleHandler := handlers.NewRulesHandler(projectRepo, endpointRepo, netCatRepo, pageSpeedRepo, pingRepo, traceRouteRepo, dataCenterRepo, taskPusher, agentHandler)
		controllers := handlers.NewHttpControllers(
			ruleHandler,
			accountRepo,
			projectRepo,
			dataCenterRepo,
			aggregateRepo,
			packageRepo,
			endpointReportRepo,
			netCatReportRepo,
			pageSPeedReportRepo,
			pingReportRepo,
			traceRouteReportRepo)

		e.GET("/", controllers.Hello)
		e.POST("/rules/register", controllers.RegisterRules, handlers.WithAuth())
		e.GET("/rules/:project_id", controllers.GetRules, handlers.WithAuth())
		e.POST("/report/endpoint/:project_id", controllers.ReportEndpoint, handlers.WithAuth())
		e.POST("/report/net_cat/:project_id", controllers.ReportNetCat, handlers.WithAuth())
		e.POST("/report/page_speed/:project_id", controllers.ReportPageSpeed, handlers.WithAuth())
		e.POST("/report/ping/:project_id", controllers.ReportPing, handlers.WithAuth())
		e.POST("/report/trace_route/:project_id", controllers.ReportTraceRoute, handlers.WithAuth())

		e.GET("/accounts/:account_id", controllers.GetAccount, handlers.WithAuth())
		e.PUT("/accounts/:account_id", controllers.UpdateAccount, handlers.WithAuth())
		e.POST("/projects", controllers.CreateProject, handlers.WithAuth())
		e.GET("/projects/:project_id", controllers.GetProject, handlers.WithAuth())
		e.PUT("/projects/:project_id", controllers.UpdateProject, handlers.WithAuth())
		e.POST("/packages", controllers.CreatePackage, handlers.WithAuth())
		e.GET("/packages/:package_id", controllers.GetPackage, handlers.WithAuth())
		e.PUT("/packages/:package_id", controllers.UpdatePackage, handlers.WithAuth())
		e.POST("/datacenters", controllers.CreateDatacenter, handlers.WithAuth())
		e.GET("/datacenters/:datacenter_id", controllers.GetDatacenter, handlers.WithAuth())
		e.PUT("/datacenters/:datacenter_id", controllers.UpdateDatacenter, handlers.WithAuth())

		e.POST("/register", controllers.Register)
		e.POST("/auth", controllers.Auth)
		e.GET("/auth/info", controllers.AuthInfo)

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
