package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"test-manager/handlers"
	"test-manager/usecase_models"
)

type EndpointTaskHandler struct {
	EndpointHandler handlers.EndpointHandler

	Logger *zap.SugaredLogger
}

func (c *EndpointTaskHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload usecase_models.Endpoints
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed on endpoint task: %v: %w", err, asynq.SkipRetry)
	}

	err := c.EndpointHandler.ExecuteEndpointRule(ctx, payload)
	if err != nil {
		c.Logger.Info(err)
		return fmt.Errorf("executing rule on endpoint task: %v", err)
	}

	c.Logger.Info("success on processing endpoint task")
	return nil
}

func NewEndpointTaskHandler(
	EndpointHandler handlers.EndpointHandler,
	Logger *zap.SugaredLogger,
) *EndpointTaskHandler {
	return &EndpointTaskHandler{
		EndpointHandler: EndpointHandler,
		Logger:          Logger,
	}
}
