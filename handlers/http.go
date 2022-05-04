package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"test-manager/usecase_models"
)

type HttpControllers interface {
	Hello(ctx echo.Context) error
	RegisterRules(ctx echo.Context) error
}

type httpControllers struct {
	rulesHandler RulesHandler
}

func NewHttpControllers(rulesHandler RulesHandler) HttpControllers {
	return &httpControllers{rulesHandler: rulesHandler}
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
