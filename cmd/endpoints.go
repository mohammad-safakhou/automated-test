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
	"test-manager/tasks/push"
	"test-manager/utils"
	"time"
)

func init() {
	rootCmd.AddCommand(endpointCmd)
}

var endpointCmd = &cobra.Command{
	Use:   "endpoint",
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

		endpointRepo := repos.NewEndpointRepository(psqlDb)
		dataCenterRepo := repos.NewDataCentersRepositoryRepository(psqlDb)

		agentHandler := handlers.NewAgentHandler()
		endpointHandler := handlers.NewEndpointHandler(endpointRepo, dataCenterRepo, taskPusher, agentHandler)
		controllers := handlers.NewHttpControllers(endpointHandler)

		e.GET("/", controllers.Hello)
		e.POST("/endpoint/register/:project_id", controllers.RegisterEndpoints)

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
