package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"test-manager/repos/influx"
	"test-manager/usecase_models"
)

type HttpControllers interface {
	Hello(ctx echo.Context) error
	RegisterRules(ctx echo.Context) error
	ReportEndpoint(ctx echo.Context) error
	ReportNetCat(ctx echo.Context) error
	ReportPageSpeed(ctx echo.Context) error
	ReportPing(ctx echo.Context) error
	ReportTraceRoute(ctx echo.Context) error
}

type httpControllers struct {
	rulesHandler               RulesHandler
	endpointReportRepository   influx.EndpointReportRepository
	netCatsReportRepository    influx.NetCatsReportRepository
	pageSpeedReportRepository  influx.PageSpeedReportRepository
	pingReportRepository       influx.PingReportRepository
	traceRouteReportRepository influx.TraceRouteReportRepository
}

func NewHttpControllers(rulesHandler RulesHandler,
	endpointReportRepository influx.EndpointReportRepository,
	netCatsReportRepository influx.NetCatsReportRepository,
	pageSpeedReportRepository influx.PageSpeedReportRepository,
	pingReportRepository influx.PingReportRepository,
	traceRouteReportRepository influx.TraceRouteReportRepository) HttpControllers {
	return &httpControllers{
		rulesHandler:               rulesHandler,
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
