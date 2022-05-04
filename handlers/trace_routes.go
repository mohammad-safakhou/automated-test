package handlers

import (
	"context"
	"github.com/labstack/gommon/log"
	"sync"
	"test-manager/repos"
	"test-manager/tasks/push"
	"test-manager/usecase_models"
)

type TraceRouteHandler interface {
	ExecuteTraceRouteRule(ctx context.Context, TraceRouteRules usecase_models.TraceRoutes) error
}

type traceRouteHandler struct {
	traceRouteRepo  repos.TraceRouteRepository
	dataCentersRepo repos.DataCentersRepository
	taskPusher      push.TaskPusher
	agentHandler    AgentHandler
}

func NewTraceRouteHandler(traceRouteRepo repos.TraceRouteRepository, dataCentersRepo repos.DataCentersRepository, taskPusher push.TaskPusher, agentHandler AgentHandler) TraceRouteHandler {
	return &traceRouteHandler{traceRouteRepo: traceRouteRepo, dataCentersRepo: dataCentersRepo, taskPusher: taskPusher, agentHandler: agentHandler}
}

func (e *traceRouteHandler) ExecuteTraceRouteRule(ctx context.Context, traceRouteRules usecase_models.TraceRoutes) error {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(traceRouteRules.Scheduling.DataCentersIds))
	for _, dataC := range traceRouteRules.Scheduling.DataCentersIds {
		go func(dataCenter int) {
			for _, rule := range traceRouteRules.TraceRouts {
				dataCenter, err := e.dataCentersRepo.GetDataCenters(ctx, dataCenter)
				if err != nil {
					log.Info("error on getting data center in executing trace rote rule: ", err)
					waitGroup.Done()
					return
				}

				response, err := e.agentHandler.SendTraceRoute(ctx, dataCenter.Baseurl, usecase_models.AgentTraceRouteRequest{
					Address: rule.Address,
					Retry:   rule.Retry,
					Hop:     rule.Hop,
				})
				if err != nil {
					log.Info("error on sending trace route in executing rule: ", err)
					waitGroup.Done()
					return
				}

				if response.Status == 0 {
					// TODO: send alert
					break
				}
			}
			waitGroup.Done()
		}(dataC)
	}

	waitGroup.Wait()
	return nil
}
