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

type NetCatHandler interface {
	ExecuteNetCatRule(ctx context.Context, NetCatRules usecase_models.NetCats) error
}

type netCatHandler struct {
	netCatRepo       repos.NetCatRepository
	dataCentersRepo  repos.DataCentersRepository
	netCatReportRepo influx.NetCatsReportRepository
	taskPusher       push.TaskPusher
	agentHandler     AgentHandler
}

func NewNetCatHandler(netCatRepo repos.NetCatRepository, dataCentersRepo repos.DataCentersRepository,netCatReportRepo influx.NetCatsReportRepository, taskPusher push.TaskPusher, agentHandler AgentHandler) NetCatHandler {
	return &netCatHandler{netCatRepo: netCatRepo, dataCentersRepo: dataCentersRepo, netCatReportRepo: netCatReportRepo,taskPusher: taskPusher, agentHandler: agentHandler}
}

func (e *netCatHandler) ExecuteNetCatRule(ctx context.Context, netCatRules usecase_models.NetCats) error {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(netCatRules.Scheduling.DataCentersIds))
	for _, dataC := range netCatRules.Scheduling.DataCentersIds {
		go func(dataCenter int) {
			for _, rule := range netCatRules.NetCats {
				dataCenter, err := e.dataCentersRepo.GetDataCenter(ctx, dataCenter)
				if err != nil {
					log.Info("error on getting data center in executing net cat rule: ", err)
					waitGroup.Done()
					return
				}

				response, err := e.agentHandler.SendNetCat(ctx, dataCenter.Baseurl, usecase_models.AgentNetCatRequest{
					Address: rule.Address,
					Port:    rule.Port,
					Type:    rule.Type,
					TimeOut: rule.TimeOut,
				})
				if err != nil {
					log.Info("error on sending net cat in executing rule: ", err)
					waitGroup.Done()
					return
				}

				if response.Status == 0 {
					err = e.netCatReportRepo.WriteNetCatsReport(ctx, netCatRules.Scheduling.ProjectId, rule.Address, 0)
					if err != nil {
						log.Info("error on writing curl report in executing rule: ", err)
					}
					// TODO: send alert
					break
				}
				err = e.netCatReportRepo.WriteNetCatsReport(ctx, netCatRules.Scheduling.ProjectId, rule.Address, 1)
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
