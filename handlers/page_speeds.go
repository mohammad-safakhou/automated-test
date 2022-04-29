package handlers

import (
	"context"
	"github.com/labstack/gommon/log"
	"sync"
	"test-manager/repos"
	"test-manager/tasks/push"
	"test-manager/usecase_models"
)

type PageSpeedHandler interface {
	ExecutePageSpeedRule(ctx context.Context, PageSpeedRules usecase_models.PageSpeed) error
}

type pageSpeedHandler struct {
	pageSpeedRepo   repos.PageSpeedRepository
	dataCentersRepo repos.DataCentersRepository
	taskPusher      push.TaskPusher
	agentHandler    AgentHandler
}

func NewPageSpeedHandler(pageSpeedRepo repos.PageSpeedRepository, dataCentersRepo repos.DataCentersRepository, taskPusher push.TaskPusher, agentHandler AgentHandler) PageSpeedHandler {
	return &pageSpeedHandler{pageSpeedRepo: pageSpeedRepo, dataCentersRepo: dataCentersRepo, taskPusher: taskPusher, agentHandler: agentHandler}
}

func (e *pageSpeedHandler) ExecutePageSpeedRule(ctx context.Context, pageSpeedRules usecase_models.PageSpeed) error {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(pageSpeedRules.Scheduling.DataCentersIds))
	for _, dataC := range pageSpeedRules.Scheduling.DataCentersIds {
		go func(dataCenter int) {
			for _, rule := range pageSpeedRules.PageSpeed {
				dataCenter, err := e.dataCentersRepo.GetDataCenters(ctx, dataCenter)
				if err != nil {
					log.Info("error on getting data center in executing page speed rule: ", err)
					waitGroup.Done()
					return
				}

				response, err := e.agentHandler.SendPageSpeed(ctx, dataCenter.Baseurl, usecase_models.AgentPageSpeedRequest{
					Url: rule.Url,
				})
				if err != nil {
					log.Info("error on sending page speed in executing rule: ", err)
					waitGroup.Done()
					return
				}

				if response.Status == 0 {
					// TODO: send alert
					break
				}
			}
		}(dataC)
	}

	waitGroup.Wait()
	return nil
}
