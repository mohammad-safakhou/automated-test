package push

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
	"test-manager/tasks/task_models"
	"test-manager/usecase_models"
	"time"
)

type TaskPusher interface {
	PushRules(ctx context.Context, payload usecase_models.RulesRequest) (taskId string, err error)
}

type taskPush struct {
	taskClient *asynq.Client
}

func NewTaskPush(taskClient *asynq.Client) TaskPusher {
	return &taskPush{taskClient: taskClient}
}

func (t *taskPush) PushRules(ctx context.Context, payload usecase_models.RulesRequest) (taskId string, err error) {
	if payload.Endpoints.Scheduling.ProjectId != 0 {
		payloadBytes, err := json.Marshal(payload.Endpoints)
		if err != nil {
			return "", err
		}

		endAt, err := time.Parse("2006-01-02 15:04:05", payload.Endpoints.Scheduling.EndAt)
		if err != nil {

			return "", err
		}
		repeat := int(endAt.Sub(time.Now()).Minutes() / float64(payload.Endpoints.Scheduling.Duration))

		for i := 0; i < repeat; i++ {
			task := asynq.NewTask(task_models.TypeEndpoint, payloadBytes)

			_, err := t.taskClient.Enqueue(
				task,
				asynq.ProcessIn(time.Duration(i)*time.Minute),
				asynq.Queue(task_models.QueueEndpoint))
			if err != nil {
				log.Println("error at enqueue endpoint task: ", err)
			}
		}
	}

	if payload.NetCats.Scheduling.ProjectId != 0 {
		payloadBytes, err := json.Marshal(payload.NetCats)
		if err != nil {
			return "", err
		}

		endAt, err := time.Parse("2006-01-02 15:04:05", payload.NetCats.Scheduling.EndAt)
		if err != nil {

			return "", err
		}
		repeat := int(endAt.Sub(time.Now()).Minutes() / float64(payload.NetCats.Scheduling.Duration))

		for i := 0; i < repeat; i++ {
			task := asynq.NewTask(task_models.TypeNetCats, payloadBytes)

			_, err := t.taskClient.Enqueue(
				task,
				asynq.ProcessIn(time.Duration(i)*time.Minute),
				asynq.Queue(task_models.QueueNetCats))
			if err != nil {
				log.Println("error at enqueue net cats task: ", err)
			}
		}
	}

	if payload.PageSpeed.Scheduling.ProjectId != 0 {
		payloadBytes, err := json.Marshal(payload.PageSpeed)
		if err != nil {
			return "", err
		}

		endAt, err := time.Parse("2006-01-02 15:04:05", payload.PageSpeed.Scheduling.EndAt)
		if err != nil {

			return "", err
		}
		repeat := int(endAt.Sub(time.Now()).Minutes() / float64(payload.PageSpeed.Scheduling.Duration))

		for i := 0; i < repeat; i++ {
			task := asynq.NewTask(task_models.TypePageSpeeds, payloadBytes)

			_, err := t.taskClient.Enqueue(
				task,
				asynq.ProcessIn(time.Duration(i)*time.Minute),
				asynq.Queue(task_models.QueuePageSpeeds))
			if err != nil {
				log.Println("error at enqueue page speeds task: ", err)
			}
		}
	}

	if payload.Pings.Scheduling.ProjectId != 0 {
		payloadBytes, err := json.Marshal(payload.Pings)
		if err != nil {
			return "", err
		}

		endAt, err := time.Parse("2006-01-02 15:04:05", payload.Pings.Scheduling.EndAt)
		if err != nil {

			return "", err
		}
		repeat := int(endAt.Sub(time.Now()).Minutes() / float64(payload.Pings.Scheduling.Duration))

		for i := 0; i < repeat; i++ {
			task := asynq.NewTask(task_models.TypePings, payloadBytes)

			_, err := t.taskClient.Enqueue(
				task,
				asynq.ProcessIn(time.Duration(i)*time.Minute),
				asynq.Queue(task_models.QueuePings))
			if err != nil {
				log.Println("error at enqueue pings task: ", err)
			}
		}
	}

	if payload.TraceRoutes.Scheduling.ProjectId != 0 {
		payloadBytes, err := json.Marshal(payload.TraceRoutes)
		if err != nil {
			return "", err
		}

		endAt, err := time.Parse("2006-01-02 15:04:05", payload.TraceRoutes.Scheduling.EndAt)
		if err != nil {

			return "", err
		}
		repeat := int(endAt.Sub(time.Now()).Minutes() / float64(payload.TraceRoutes.Scheduling.Duration))

		for i := 0; i < repeat; i++ {
			task := asynq.NewTask(task_models.TypeTraceRoutes, payloadBytes)

			_, err := t.taskClient.Enqueue(
				task,
				asynq.ProcessIn(time.Duration(i)*time.Minute),
				asynq.Queue(task_models.QueueTraceRoutes))
			if err != nil {
				log.Println("error at enqueue trace routes task: ", err)
			}
		}
	}

	return "", nil
}
