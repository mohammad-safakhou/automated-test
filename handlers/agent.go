package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"test-manager/usecase_models"
)

type AgentHandler interface {
	SendCurl(ctx context.Context, dataCenterUrl string, request usecase_models.AgentCurlRequest) (response []byte, responseHeader map[string][]string, status string, err error)
}

type agentHandler struct {
}

func NewAgentHandler() AgentHandler {
	return &agentHandler{}
}

func (a *agentHandler) SendCurl(ctx context.Context, dataCenterUrl string, request usecase_models.AgentCurlRequest) (response []byte, responseHeader map[string][]string, status string, err error) {
	reqB, _ := json.Marshal(request)
	req, err := http.NewRequestWithContext(ctx, "POST", dataCenterUrl+"/v1/curl", bytes.NewBuffer(reqB))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return response, responseHeader, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	return respBody, resp.Header, resp.Status, nil
}
