package handlers

import (
	"context"
	"github.com/labstack/gommon/log"
	"sync"
	"test-manager/repos"
	"test-manager/repos/influx"
	"test-manager/tasks/push"
	"test-manager/usecase_models"
)

type PageSpeedHandler interface {
	ExecutePageSpeedRule(ctx context.Context, PageSpeedRules usecase_models.PageSpeed) error
}

type pageSpeedHandler struct {
	pageSpeedRepo       repos.PageSpeedRepository
	dataCentersRepo     repos.DataCentersRepository
	pageSpeedReportRepo influx.PageSpeedReportRepository
	taskPusher          push.TaskPusher
	agentHandler        AgentHandler
}

func NewPageSpeedHandler(pageSpeedRepo repos.PageSpeedRepository, dataCentersRepo repos.DataCentersRepository, pageSpeedReportRepo influx.PageSpeedReportRepository, taskPusher push.TaskPusher, agentHandler AgentHandler) PageSpeedHandler {
	return &pageSpeedHandler{pageSpeedRepo: pageSpeedRepo, dataCentersRepo: dataCentersRepo, pageSpeedReportRepo: pageSpeedReportRepo, taskPusher: taskPusher, agentHandler: agentHandler}
}

func (e *pageSpeedHandler) ExecutePageSpeedRule(ctx context.Context, pageSpeedRules usecase_models.PageSpeed) error {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(pageSpeedRules.Scheduling.DataCentersIds))
	for _, dataC := range pageSpeedRules.Scheduling.DataCentersIds {
		go func(dataCenter int) {
			for _, rule := range pageSpeedRules.PageSpeed {
				dataCenter, err := e.dataCentersRepo.GetDataCenter(ctx, dataCenter)
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
					err = e.pageSpeedReportRepo.WritePageSpeedReport(ctx, pageSpeedRules.Scheduling.ProjectId, rule.Url, 0)
					if err != nil {
						log.Info("error on writing curl report in executing rule: ", err)
					}
					// TODO: send alert
					break
				}
				err = e.pageSpeedReportRepo.WritePageSpeedReport(ctx, pageSpeedRules.Scheduling.ProjectId, rule.Url, 1)
				if err != nil {
					log.Info("error on writing curl report in executing rule: ", err)
				}
			}
			waitGroup.Done()
		}(dataC)
	}

	waitGroup.Wait()
	return nil
}
