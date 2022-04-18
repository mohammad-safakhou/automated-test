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
	PushToEndpoint(ctx context.Context, payload usecase_models.EndpointRequest) (taskId string, err error)
}

type taskPush struct {
	taskClient *asynq.Client
}

func NewTaskPush(taskClient *asynq.Client) TaskPusher {
	return &taskPush{taskClient: taskClient}
}

func (t *taskPush) PushToEndpoint(ctx context.Context, payload usecase_models.EndpointRequest) (taskId string, err error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	endAt, err := time.Parse("2006-01-02 15:04:05", payload.Scheduling.EndAt)
	if err != nil {

		return "", err
	}
	repeat := int(endAt.Sub(time.Now()).Minutes() / float64(payload.Scheduling.Duration))

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

	return "", nil
}
