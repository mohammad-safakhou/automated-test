package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/volatiletech/null/v8"
	"test-manager/repos"
	"test-manager/tasks/push"
	"test-manager/usecase_models"
	models "test-manager/usecase_models/boiler"
)

type RulesHandler interface {
	RegisterRules(ctx context.Context, rules usecase_models.RulesRequest) error
}

type rulesHandler struct {
	projectRepo     repos.ProjectsRepository
	endpointRepo    repos.EndpointRepository
	netCatRepo      repos.NetCatRepository
	pageSpeedRepo   repos.PageSpeedRepository
	pingRepo        repos.PingRepository
	traceRouteRepo  repos.TraceRouteRepository
	dataCentersRepo repos.DataCentersRepository
	taskPusher      push.TaskPusher
	agentHandler    AgentHandler
}

func NewRulesHandler(
	projectRepo repos.ProjectsRepository,
	endpointRepo repos.EndpointRepository,
	netCatRepo repos.NetCatRepository,
	pageSpeedRepo repos.PageSpeedRepository,
	pingRepo repos.PingRepository,
	traceRouteRepo repos.TraceRouteRepository,
	dataCentersRepo repos.DataCentersRepository,
	taskPusher push.TaskPusher,
	agentHandler AgentHandler,
) RulesHandler {
	return &rulesHandler{
		projectRepo:     projectRepo,
		endpointRepo:    endpointRepo,
		netCatRepo:      netCatRepo,
		pageSpeedRepo:   pageSpeedRepo,
		pingRepo:        pingRepo,
		traceRouteRepo:  traceRouteRepo,
		dataCentersRepo: dataCentersRepo,
		taskPusher:      taskPusher,
		agentHandler:    agentHandler,
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (r *rulesHandler) RegisterRules(ctx context.Context, rules usecase_models.RulesRequest) error {
	projects, err := r.projectRepo.GetProjects(ctx, IdentityStruct.Id)
	if err != nil {
		return err
	}
	var projectIds []int
	for _, value := range projects {
		projectIds = append(projectIds, value.ID)
	}
	if !contains(projectIds, rules.Endpoints.Scheduling.ProjectId) {
		return errors.New(fmt.Sprintf("project id : %d is not your project", rules.Endpoints.Scheduling.ProjectId))
	}
	if !contains(projectIds, rules.NetCats.Scheduling.ProjectId) {
		return errors.New(fmt.Sprintf("project id : %d is not your project", rules.NetCats.Scheduling.ProjectId))
	}
	if !contains(projectIds, rules.Pings.Scheduling.ProjectId) {
		return errors.New(fmt.Sprintf("project id : %d is not your project", rules.Pings.Scheduling.ProjectId))
	}
	if !contains(projectIds, rules.TraceRoutes.Scheduling.ProjectId) {
		return errors.New(fmt.Sprintf("project id : %d is not your project", rules.TraceRoutes.Scheduling.ProjectId))
	}
	if !contains(projectIds, rules.PageSpeed.Scheduling.ProjectId) {
		return errors.New(fmt.Sprintf("project id : %d is not your project", rules.PageSpeed.Scheduling.ProjectId))
	}

	if len(rules.Endpoints.Endpoints) != 0 {
		j, _ := json.Marshal(rules.Endpoints)
		rulesStr := string(j)
		endpointId, err := r.endpointRepo.SaveEndpoint(ctx, models.Endpoint{
			Data:      null.NewString(rulesStr, true),
			ProjectID: rules.Endpoints.Scheduling.ProjectId,
		})
		rules.Endpoints.Scheduling.PipelineId = endpointId
		if err != nil {
			return err
		}
	}
	if len(rules.NetCats.NetCats) != 0 {
		j, _ := json.Marshal(rules.NetCats)
		rulesStr := string(j)
		_, err := r.netCatRepo.SaveNetCat(ctx, models.NetCat{
			Data:      null.NewString(rulesStr, true),
			ProjectID: rules.NetCats.Scheduling.ProjectId,
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
			ProjectID: rules.PageSpeed.Scheduling.ProjectId,
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
			ProjectID: rules.Pings.Scheduling.ProjectId,
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
			ProjectID: rules.TraceRoutes.Scheduling.ProjectId,
		})
		if err != nil {
			return err
		}
	}

	_, err = r.taskPusher.PushRules(ctx, rules)
	if err != nil {
		return err
	}

	return nil
}
