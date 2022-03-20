package handlers

import "test-manager/models"

type EndpointHandler interface {
	RegisterRules(rules []models.EndpointRequest) error
}

type endpointHandler struct {
}

func NewEndpointHandler() EndpointHandler {
	return &endpointHandler{}
}

func (e *endpointHandler) RegisterRules(rules []models.EndpointRequest) error {
	return nil
}
