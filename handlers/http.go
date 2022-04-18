package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"test-manager/usecase_models"
)

type HttpControllers interface {
	Hello(ctx echo.Context) error
	RegisterEndpoints(ctx echo.Context) error
}

type httpControllers struct {
	endpointHandler EndpointHandler
}

func NewHttpControllers(endpointHandler EndpointHandler) HttpControllers {
	return &httpControllers{endpointHandler: endpointHandler}
}

func (hc *httpControllers) Hello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "yo")
}

func (hc *httpControllers) RegisterEndpoints(ctx echo.Context) error {
	req := new(usecase_models.EndpointRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	projectId, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = hc.endpointHandler.RegisterRules(context.TODO(), *req, projectId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "ok")
}
