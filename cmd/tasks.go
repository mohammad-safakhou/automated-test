package cmd

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/cobra"
	"math/rand"
	"test-manager/handlers"
	"test-manager/repos"
	"test-manager/repos/influx"
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
		task_models.QueueEndpoint:    6,
		task_models.QueueNetCats:     6,
		task_models.QueuePageSpeeds:  6,
		task_models.QueuePings:       6,
		task_models.QueueTraceRoutes: 6,
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
		influxClient, writeAPI, queryAPI, err := utils.CreateInfluxDBConnection(context.TODO(), "GNDVtSQxQ_weUoyLpsWQIl_PK62ugeFJxQ2KbOP-lJZ5SRpu2cuQmkP-QQF78b_EfrI_mWrg5kxnNDDUCnMb6A==", "http://localhost:8086", "test", "my-bucket")
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

		for {
			st := 0
			x := rand.Intn(500)
			if x >= 250 {
				st = 1
			} else {
				st = 0
			}
			p := influxdb2.NewPoint("endpoint",
				map[string]string{"project_id": "1", "pipeline_id": "1"},
				map[string]interface{}{"success": st, "response_time": x},
				time.Now())
			err = writeAPI.WritePoint(context.Background(), p)
			if err != nil {
				fmt.Println(err.Error())
			}
			time.Sleep(5 * time.Second)
		}
		go func() {
			for {
				err, res := endpointReportRepo.ReadEndpointReportByProject(context.TODO(), 1, "1h")
				fmt.Println(res)
				fmt.Println(err)
				time.Sleep(5 * time.Second)
			}
		}()
		agentHandler := handlers.NewAgentHandler()
		endpointHandler := handlers.NewEndpointHandler(endpointRepo, dataCenterRepo, endpointReportRepo, taskPusher, agentHandler)
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
