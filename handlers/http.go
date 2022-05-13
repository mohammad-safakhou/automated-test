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
}

type httpControllers struct {
	rulesHandler             RulesHandler
	endpointReportRepository influx.EndpointReportRepository
}

func NewHttpControllers(rulesHandler RulesHandler, endpointReportRepository influx.EndpointReportRepository) HttpControllers {
	return &httpControllers{rulesHandler: rulesHandler, endpointReportRepository: endpointReportRepository}
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
