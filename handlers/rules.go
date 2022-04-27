package handlers

import (
	"context"
	"encoding/json"
	"github.com/volatiletech/null/v8"
	"test-manager/repos"
	"test-manager/tasks/push"
	"test-manager/usecase_models"
	models "test-manager/usecase_models/boiler"
)

type RulesHandler interface {
	RegisterRules(ctx context.Context, rules usecase_models.RulesRequest, projectId int) error
}

type rulesHandler struct {
	endpointRepo    repos.EndpointRepository
	netCatRepo      repos.NetCatRepository
	pageSpeedRepo   repos.PageSpeedRepository
	pingRepo        repos.PingRepository
	traceRouteRepo  repos.TraceRouteRepository
	dataCentersRepo repos.DataCentersRepository
	taskPusher      push.TaskPusher
	agentHandler    AgentHandler
}

func NewRulesHandler() RulesHandler {
	return &rulesHandler{}
}

func (r *rulesHandler) RegisterRules(ctx context.Context, rules usecase_models.RulesRequest, projectId int) error {
	if len(rules.Endpoints.Endpoints) != 0 {
		j, _ := json.Marshal(rules.Endpoints)
		rulesStr := string(j)
		_, err := r.endpointRepo.SaveEndpoint(ctx, models.Endpoint{
			Data:      null.NewString(rulesStr, true),
			ProjectID: projectId,
		})
		if err != nil {
			return err
		}
	}
	if len(rules.NetCats.NetCats) != 0 {
		j, _ := json.Marshal(rules.NetCats)
		rulesStr := string(j)
		_, err := r.netCatRepo.SaveNetCat(ctx, models.NetCat{
			Data:      null.NewString(rulesStr, true),
			ProjectID: projectId,
		})
		if err != nil {
			return err
		}
	}
	if len(rules.PageSpeed.PageSpeed) != 0 {
		j, _ := json.Marshal(rules.PageSpeed)
		rulesStr := string(j)
		_, err := r.pageSpeedRepo.SavePageSpeed(ctx, models.PageSpeed{
			Data:      null.NewString(rulesStr, true),
			ProjectID: projectId,
		})
		if err != nil {
			return err
		}
	}
	if len(rules.Pings.Pings) != 0 {
		j, _ := json.Marshal(rules.Pings)
		rulesStr := string(j)
		_, err := r.pingRepo.SavePing(ctx, models.Ping{
			Data:      null.NewString(rulesStr, true),
			ProjectID: projectId,
		})
		if err != nil {
			return err
		}
	}
	if len(rules.TraceRoutes.TraceRouts) != 0 {
		j, _ := json.Marshal(rules.TraceRoutes)
		rulesStr := string(j)
		_, err := r.traceRouteRepo.SaveTraceRoute(ctx, models.TraceRoute{
			Data:      null.NewString(rulesStr, true),
			ProjectID: projectId,
		})
		if err != nil {
			return err
		}
	}

	_, err := r.taskPusher.PushRules(ctx, rules)
	if err != nil {
		return err
	}

	return nil
}
