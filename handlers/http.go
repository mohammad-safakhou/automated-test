package handlers

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"net/http"
	"strconv"
	"test-manager/repos"
	"test-manager/repos/influx"
	"test-manager/usecase_models"
	models "test-manager/usecase_models/boiler"
)

type HttpControllers interface {
	Hello(ctx echo.Context) error
	RegisterRules(ctx echo.Context) error
	GetRules(ctx echo.Context) error
	ReportEndpoint(ctx echo.Context) error
	ReportNetCat(ctx echo.Context) error
	ReportPageSpeed(ctx echo.Context) error
	ReportPing(ctx echo.Context) error
	ReportTraceRoute(ctx echo.Context) error

	GetAccount(ctx echo.Context) error
	UpdateAccount(ctx echo.Context) error
	CreateProject(ctx echo.Context) error
	GetProject(ctx echo.Context) error
	UpdateProject(ctx echo.Context) error
	CreateDatacenter(ctx echo.Context) error
	GetDatacenter(ctx echo.Context) error
	UpdateDatacenter(ctx echo.Context) error

	Register(ctx echo.Context) error
	Auth(ctx echo.Context) error
	AuthInfo(ctx echo.Context) error
}

type httpControllers struct {
	rulesHandler               RulesHandler
	accountRepo                repos.AccountsRepository
	projectRepo                repos.ProjectsRepository
	datacenterRepo             repos.DataCentersRepository
	aggregateRepository        repos.AggregateRepository
	endpointReportRepository   influx.EndpointReportRepository
	netCatsReportRepository    influx.NetCatsReportRepository
	pageSpeedReportRepository  influx.PageSpeedReportRepository
	pingReportRepository       influx.PingReportRepository
	traceRouteReportRepository influx.TraceRouteReportRepository
}

func NewHttpControllers(rulesHandler RulesHandler,
	accountRepo repos.AccountsRepository,
	projectRepo repos.ProjectsRepository,
	aggregateRepository repos.AggregateRepository,
	endpointReportRepository influx.EndpointReportRepository,
	netCatsReportRepository influx.NetCatsReportRepository,
	pageSpeedReportRepository influx.PageSpeedReportRepository,
	pingReportRepository influx.PingReportRepository,
	traceRouteReportRepository influx.TraceRouteReportRepository) HttpControllers {
	return &httpControllers{
		rulesHandler:               rulesHandler,
		accountRepo:                accountRepo,
		projectRepo:                projectRepo,
		aggregateRepository:        aggregateRepository,
		endpointReportRepository:   endpointReportRepository,
		netCatsReportRepository:    netCatsReportRepository,
		pageSpeedReportRepository:  pageSpeedReportRepository,
		pingReportRepository:       pingReportRepository,
		traceRouteReportRepository: traceRouteReportRepository,
	}
}

func (hc *httpControllers) Hello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "yo")
}

func (hc *httpControllers) GetAccount(ctx echo.Context) error {
	accountId, err := strconv.Atoi(ctx.Param("account_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	account, err := hc.accountRepo.GetAccounts(ctx.Request().Context(), accountId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, usecase_models.AccountResponse{
		ID:          accountId,
		FirstName:   account.FirstName.String,
		LastName:    account.LastName.String,
		PhoneNumber: account.PhoneNumber.String,
		Email:       account.Email.String,
		Username:    account.Username.String,
	})
}

func (hc *httpControllers) UpdateAccount(ctx echo.Context) error {
	req := new(usecase_models.Account)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if req.Password != "" {
		plainText, err := PrivateKey.Decrypt(nil, []byte(req.Password), &rsa.OAEPOptions{Hash: crypto.SHA256})
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		req.Password = string(plainText)
	}

	err := hc.accountRepo.UpdateAccounts(ctx.Request().Context(), models.Account{
		FirstName:   null.NewString(req.FirstName, true),
		LastName:    null.NewString(req.LastName, true),
		PhoneNumber: null.NewString(req.PhoneNumber, true),
		Email:       null.NewString(req.Email, true),
		Username:    null.NewString(req.Username, true),
		Password:    null.NewString(req.Password, true),
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, "")
}

func (hc *httpControllers) CreateProject(ctx echo.Context) error {
	req := new(usecase_models.Project)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	notif, err := json.Marshal(req.Notifications)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	projectId, err := hc.projectRepo.SaveProjects(ctx.Request().Context(), models.Project{
		Title:         null.NewString(req.Title, true),
		IsActive:      null.NewBool(req.IsActive, true),
		ExpireAt:      null.NewTime(req.ExpireAt, true),
		AccountID:     IdentityStruct.Id,
		Notifications: null.NewJSON(notif, true),
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, usecase_models.CreateProjectResponse{
		ProjectId: projectId,
	})
}

func (hc *httpControllers) GetProject(ctx echo.Context) error {
	projectId := 0
	var err error
	if ctx.Param("project_id") != "" {
		projectId, err = strconv.Atoi(ctx.Param("project_id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
	}

	var projects []*models.Project
	if projectId != 0 {
		project, err := hc.projectRepo.GetProject(ctx.Request().Context(), IdentityStruct.Id, projectId)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		projects = append(projects, &project)
	} else {
		projects, err = hc.projectRepo.GetProjects(ctx.Request().Context(), IdentityStruct.Id)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		projects = append(projects, projects...)
	}

	var projectsResponse []usecase_models.Project
	for _, project := range projects {
		var notifications usecase_models.Notifications
		err = json.Unmarshal(project.Notifications.JSON, &notifications)
		if err != nil {
			continue
		}
		projectsResponse = append(projectsResponse, usecase_models.Project{
			ID:            project.ID,
			Title:         project.Title.String,
			IsActive:      project.IsActive.Bool,
			Notifications: notifications,
			ExpireAt:      project.ExpireAt.Time,
			UpdatedAt:     project.UpdatedAt,
			CreatedAt:     project.CreatedAt,
			DeletedAt:     project.DeletedAt.Time,
		})
	}

	return ctx.JSON(http.StatusOK, projectsResponse)
}

func (hc *httpControllers) UpdateProject(ctx echo.Context) error {
	req := new(usecase_models.Project)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	projectId, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	notifications, err := json.Marshal(req.Notifications)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	err = hc.projectRepo.UpdateProjects(ctx.Request().Context(), models.Project{
		ID:            projectId,
		Title:         null.NewString(req.Title, true),
		IsActive:      null.NewBool(req.IsActive, true),
		ExpireAt:      null.NewTime(req.ExpireAt, true),
		Notifications: null.NewJSON(notifications, true),
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, "")
}

func (hc *httpControllers) CreateDatacenter(ctx echo.Context) error {
	req := new(usecase_models.Datacenter)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	datacenterId, err := hc.datacenterRepo.SaveDataCenters(ctx.Request().Context(), models.Datacenter{
		Baseurl:        req.Baseurl,
		Title:          req.Title,
		ConnectionRate: req.ConnectionRate,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, usecase_models.CreateDatacenterResponse{DatacenterId: datacenterId})
}

func (hc *httpControllers) GetDatacenter(ctx echo.Context) error {
	datacenterId := 0
	var err error
	if ctx.QueryParam("datacenter_id") != "" {
		datacenterId, err = strconv.Atoi(ctx.QueryParam("datacenter_id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
	}

	var datacenters []*models.Datacenter
	if datacenterId != 0 {
		datacenter, err := hc.datacenterRepo.GetDataCenter(ctx.Request().Context(), datacenterId)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		datacenters = append(datacenters, &datacenter)
	} else {
		datacenters, err = hc.datacenterRepo.GetDataCenters(ctx.Request().Context())
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		datacenters = append(datacenters, datacenters...)
	}

	var datacentersResponse []usecase_models.Datacenter
	for _, datacenter := range datacenters {
		datacentersResponse = append(datacentersResponse, usecase_models.Datacenter{
			ID:             datacenter.ID,
			Baseurl:        datacenter.Baseurl,
			Title:          datacenter.Title,
			ConnectionRate: datacenter.ConnectionRate,
			UpdatedAt:      datacenter.UpdatedAt,
			CreatedAt:      datacenter.CreatedAt,
			DeletedAt:      datacenter.DeletedAt,
		})
	}

	return ctx.JSON(http.StatusOK, datacentersResponse)
}

func (hc *httpControllers) UpdateDatacenter(ctx echo.Context) error {
	req := new(usecase_models.Datacenter)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	datacenterId, err := strconv.Atoi(ctx.Param("datacenter_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	err = hc.datacenterRepo.UpdateDataCenters(ctx.Request().Context(), models.Datacenter{
		ID:             datacenterId,
		Baseurl:        req.Baseurl,
		Title:          req.Title,
		ConnectionRate: req.ConnectionRate,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, "")
}

func (hc *httpControllers) Register(ctx echo.Context) error {
	req := new(usecase_models.Account)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	plainText, err := PrivateKey.Decrypt(nil, []byte(req.Password), &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	accountId, err := hc.accountRepo.SaveAccounts(ctx.Request().Context(), models.Account{
		FirstName:   null.NewString(req.FirstName, true),
		LastName:    null.NewString(req.LastName, true),
		PhoneNumber: null.NewString(req.PhoneNumber, true),
		Email:       null.NewString(req.Email, true),
		Username:    null.NewString(req.Username, true),
		Password:    null.NewString(string(plainText), true),
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, usecase_models.RegisterAccountResponse{
		AccountId: accountId,
	})
}

func (hc *httpControllers) Auth(ctx echo.Context) error {
	req := new(usecase_models.Auth)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	account, err := hc.accountRepo.GetAccounts(ctx.Request().Context(), IdentityStruct.Id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	plainText, err := PrivateKey.Decrypt(nil, []byte(req.Password), &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	if string(plainText) != account.Password.String {
		return ctx.JSON(http.StatusUnauthorized, "password or username incorrect")
	}

	token, err := NewJWTToken(jwt.StandardClaims{
		Audience: Aud,
		Id:       strconv.Itoa(account.ID),
	})
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "password or username incorrect")
	}

	return ctx.JSON(http.StatusOK, usecase_models.AuthResponse{
		Token: token,
	})
}

func (hc *httpControllers) AuthInfo(ctx echo.Context) error {
	resp := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&PrivateKey.PublicKey),
		},
	)
	return ctx.JSON(http.StatusOK, resp)
}

func (hc *httpControllers) RegisterRules(ctx echo.Context) error {
	req := new(usecase_models.RulesRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := hc.rulesHandler.RegisterRules(context.TODO(), *req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "ok")
}

func (hc *httpControllers) GetRules(ctx echo.Context) error {
	projectId, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	subWorks, err := hc.aggregateRepository.AggregateAllRuleSubWorks(ctx.Request().Context(), projectId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, subWorks)
}

type ReportEndpointModel struct {
	PipelineId int      `json:"pipeline_id"`
	TimeFrame  string   `json:"timeframe"`
	Fields     []string `json:"fields"`
}

func (hc *httpControllers) ReportEndpoint(ctx echo.Context) error {
	projectId, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	req := new(ReportEndpointModel)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err, result := hc.endpointReportRepository.ReadEndpointReportByProject(ctx.Request().Context(), projectId, req.PipelineId, req.TimeFrame, req.Fields)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, result)
}

type ReportNetCatModel struct {
	Url       string   `json:"url"`
	TimeFrame string   `json:"timeframe"`
	Fields    []string `json:"fields"`
}

func (hc *httpControllers) ReportNetCat(ctx echo.Context) error {
	projectId, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	req := new(ReportNetCatModel)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err, result := hc.netCatsReportRepository.ReadNetCatsReportByProject(ctx.Request().Context(), projectId, req.Url, req.TimeFrame, req.Fields)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, result)
}

type ReportPageSpeedModel struct {
	Url       string   `json:"url"`
	TimeFrame string   `json:"timeframe"`
	Fields    []string `json:"fields"`
}

func (hc *httpControllers) ReportPageSpeed(ctx echo.Context) error {
	projectId, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	req := new(ReportPageSpeedModel)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err, result := hc.pageSpeedReportRepository.ReadPageSpeedReportByProject(ctx.Request().Context(), projectId, req.Url, req.TimeFrame, req.Fields)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, result)
}

type ReportPingModel struct {
	Url       string   `json:"url"`
	TimeFrame string   `json:"timeframe"`
	Fields    []string `json:"fields"`
}

func (hc *httpControllers) ReportPing(ctx echo.Context) error {
	projectId, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	req := new(ReportPingModel)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err, result := hc.pingReportRepository.ReadPingReportByProject(ctx.Request().Context(), projectId, req.Url, req.TimeFrame, req.Fields)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, result)
}

type ReportTraceRouteModel struct {
	Url       string   `json:"url"`
	TimeFrame string   `json:"timeframe"`
	Fields    []string `json:"fields"`
}

func (hc *httpControllers) ReportTraceRoute(ctx echo.Context) error {
	projectId, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	req := new(ReportTraceRouteModel)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err, result := hc.traceRouteReportRepository.ReadTraceRouteReportByProject(ctx.Request().Context(), projectId, req.Url, req.TimeFrame, req.Fields)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, result)
}
