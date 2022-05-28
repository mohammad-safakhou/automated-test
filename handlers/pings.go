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

type PingHandler interface {
	ExecutePingRule(ctx context.Context, PingRules usecase_models.Pings) error
}

type pingHandler struct {
	pingRepo        repos.PingRepository
	dataCentersRepo repos.DataCentersRepository
	pingReportRepo  influx.PingReportRepository
	taskPusher      push.TaskPusher
	agentHandler    AgentHandler
}

func NewPingHandler(pingRepo repos.PingRepository, dataCentersRepo repos.DataCentersRepository, pingReportRepo influx.PingReportRepository, taskPusher push.TaskPusher, agentHandler AgentHandler) PingHandler {
	return &pingHandler{pingRepo: pingRepo, dataCentersRepo: dataCentersRepo, pingReportRepo: pingReportRepo, taskPusher: taskPusher, agentHandler: agentHandler}
}

func (e *pingHandler) ExecutePingRule(ctx context.Context, pingRules usecase_models.Pings) error {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(pingRules.Scheduling.DataCentersIds))
	for _, dataC := range pingRules.Scheduling.DataCentersIds {
		go func(dataCenter int) {
			for _, rule := range pingRules.Pings {
				dataCenter, err := e.dataCentersRepo.GetDataCenters(ctx, dataCenter)
				if err != nil {
					log.Info("error on getting data center in executing ping rule: ", err)
					waitGroup.Done()
					return
				}

				response, err := e.agentHandler.SendPing(ctx, dataCenter.Baseurl, usecase_models.AgentPingRequest{
					Address: rule.Address,
					Count:   rule.Count,
					TimeOut: rule.TimeOut,
				})
				if err != nil {
					log.Info("error on sending ping in executing rule: ", err)
					waitGroup.Done()
					return
				}

				if response.Status == 0 {
					err = e.pingReportRepo.WritePingReport(ctx, pingRules.Scheduling.ProjectId, rule.Address, 0)
					if err != nil {
						log.Info("error on writing curl report in executing rule: ", err)
					}
					// TODO: send alert
					break
				}
				err = e.pingReportRepo.WritePingReport(ctx, pingRules.Scheduling.ProjectId, rule.Address, 1)
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
