package cmd

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
	"test-manager/handlers"
	"test-manager/repos"
	"test-manager/tasks"
	"test-manager/tasks/push"
	"test-manager/tasks/task_models"
	"test-manager/utils"
	"time"
)

const (
	numWorkersAsynq = 20
)

func init() {
	rootCmd.AddCommand(consumeTasksCmd)
}

var consumeTasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "consume async tasks events",
	Long:  `consume async tasks events`,
	Run: func(cmd *cobra.Command, args []string) {
		l := utils.ZapLogger()
		zLogger := l.Sugar()

		redisClient, err := utils.CreateRedisConnection(context.TODO(), "localhost", "6379", 3*time.Second)
		if err != nil {
			panic(err)
		}

		psqlDb, err := utils.PostgresConnection("localhost", "5432", "root", "root", "tester", "disable")
		if err != nil {
			panic(err)
		}

		srv := asynq.NewServer(
			asynq.RedisClientOpt{
				Addr:        redisClient.Options().Addr,
				DialTimeout: redisClient.Options().DialTimeout,
				Username:    redisClient.Options().Username,
				Password:    redisClient.Options().Password,
			}, asynq.Config{
				Concurrency: numWorkersAsynq,
				Logger:      zLogger,
			},
		)

		asynqClient := asynq.NewClient(asynq.RedisClientOpt{
			Addr:        redisClient.Options().Addr,
			DialTimeout: redisClient.Options().DialTimeout,
			Username:    redisClient.Options().Username,
			Password:    redisClient.Options().Password,
		})
		taskPusher := push.NewTaskPush(asynqClient)

		endpointRepo := repos.NewEndpointRepository(psqlDb)

		endpointHandler := handlers.NewEndpointHandler(endpointRepo, taskPusher)

		mux := asynq.NewServeMux()
		// handlers
		mux.Handle(task_models.TypeEndpoint, tasks.NewEndpointTaskHandler(endpointHandler, zLogger))

		if err := srv.Run(mux); err != nil {
			zLogger.Fatalf("cant start server: %s", err)
		}
	},
}
