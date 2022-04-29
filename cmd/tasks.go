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

var (
	// list of queues associated with priority, large numbers indicate higher priority
	queues = map[string]int{
		task_models.QueueEndpoint: 6,
	}
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
				Queues:      queues,
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
		netCatRepo := repos.NewNetCatRepository(psqlDb)
		pageSpeedRepo := repos.NewPageSpeedRepository(psqlDb)
		pingRepo := repos.NewPingRepository(psqlDb)
		traceRouteRepo := repos.NewTraceRouteRepository(psqlDb)
		dataCenterRepo := repos.NewDataCentersRepositoryRepository(psqlDb)

		agentHandler := handlers.NewAgentHandler()
		endpointHandler := handlers.NewEndpointHandler(endpointRepo, dataCenterRepo, taskPusher, agentHandler)
		netCatHandler := handlers.NewNetCatHandler(netCatRepo, dataCenterRepo, taskPusher, agentHandler)
		pageSpeedHandler := handlers.NewPageSpeedHandler(pageSpeedRepo, dataCenterRepo, taskPusher, agentHandler)
		pingHandler := handlers.NewPingHandler(pingRepo, dataCenterRepo, taskPusher, agentHandler)
		traceRouteHandler := handlers.NewTraceRouteHandler(traceRouteRepo, dataCenterRepo, taskPusher, agentHandler)

		mux := asynq.NewServeMux()
		// handlers
		mux.Handle(task_models.TypeEndpoint, tasks.NewEndpointTaskHandler(endpointHandler, zLogger))
		mux.Handle(task_models.TypeNetCats, tasks.NewNetCatTaskHandler(netCatHandler, zLogger))
		mux.Handle(task_models.TypePageSpeeds, tasks.NewPageSpeedTaskHandler(pageSpeedHandler, zLogger))
		mux.Handle(task_models.TypePings, tasks.NewPingTaskHandler(pingHandler, zLogger))
		mux.Handle(task_models.TypeTraceRoutes, tasks.NewTraceRouteTaskHandler(traceRouteHandler, zLogger))

		if err := srv.Run(mux); err != nil {
			zLogger.Fatalf("cant start server: %s", err)
		}
	},
}
