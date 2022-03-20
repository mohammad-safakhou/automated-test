package handlers

type EndpointHandler interface {
	RegisterRules()
}

type endpointHandler struct {
}

func NewEndpointHandler() EndpointHandler {
	return &endpointHandler{}
}

func (e *endpointHandler) RegisterRules() {

}
