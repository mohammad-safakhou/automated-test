package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/tidwall/gjson"
	"strings"
	"sync"
	"test-manager/repos"
	"test-manager/repos/influx"
	"test-manager/tasks/push"
	"test-manager/usecase_models"
)

type EndpointHandler interface {
	ExecuteEndpointRule(ctx context.Context, endpointRules usecase_models.Endpoints) error
}

type endpointHandler struct {
	endpointRepo       repos.EndpointRepository
	dataCentersRepo    repos.DataCentersRepository
	endpointReportRepo influx.EndpointReportRepository
	taskPusher         push.TaskPusher
	agentHandler       AgentHandler
}

func NewEndpointHandler(endpointRepo repos.EndpointRepository, dataCentersRepo repos.DataCentersRepository, endpointReportRepo influx.EndpointReportRepository, taskPusher push.TaskPusher, agentHandler AgentHandler) EndpointHandler {
	return &endpointHandler{endpointRepo: endpointRepo, dataCentersRepo: dataCentersRepo, endpointReportRepo: endpointReportRepo, taskPusher: taskPusher, agentHandler: agentHandler}
}

func (e *endpointHandler) ExecuteEndpointRule(ctx context.Context, endpointRules usecase_models.Endpoints) error {
	var responses = usecase_models.EndpointResponses{
		HeaderResponses: map[int]map[string][]string{},
		BodyResponses:   map[int]string{},
	}

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(endpointRules.Scheduling.DataCentersIds))
	for _, dataC := range endpointRules.Scheduling.DataCentersIds {
		go func(dataCenter int) {
			for _, rule := range endpointRules.Endpoints {
				var value []string
				for _, bodyDependency := range rule.BodyDependency {
					if len(bodyDependency.ParentKey) >= 8 && bodyDependency.ParentKey[0:8] == "$header_" {
						value = responses.HeaderResponses[bodyDependency.EndpointId][bodyDependency.ParentKey[8:]]
					} else if len(bodyDependency.ParentKey) >= 6 && bodyDependency.ParentKey[0:6] == "$body_" {
						value = []string{gjson.Get(responses.BodyResponses[bodyDependency.EndpointId], bodyDependency.ParentKey[6:]).String()}
					} else {
						log.Info("error on checking body dependency in executing rule")
						waitGroup.Done()
						return
					}
					ruleBody := gjson.Parse(rule.Body).Value().(map[string]interface{})
					ruleBody[bodyDependency.Key] = strings.Join(value[:], ",")
					newBody, err := json.Marshal(ruleBody)
					if err != nil {
						log.Info("error on json marshal rule body in executing rule: ", err)
						waitGroup.Done()
						return
					}
					rule.Body = string(newBody)
				}
				for _, headerDependency := range rule.HeaderDependency {
					var value []string
					if len(headerDependency.ParentKey) >= 8 && headerDependency.ParentKey[0:8] == "$header_" {
						value = responses.HeaderResponses[headerDependency.EndpointId][headerDependency.ParentKey[8:]]
					} else if len(headerDependency.ParentKey) >= 6 && headerDependency.ParentKey[0:6] == "$body_" {
						value = []string{gjson.Get(responses.BodyResponses[headerDependency.EndpointId], headerDependency.ParentKey[6:]).String()}
					} else {
						log.Info("error on checking header dependency in executing rule")
						waitGroup.Done()
						return
					}
					rule.Header[headerDependency.Key] = strings.Join(value[:], ",")
				}

				dataCenter, err := e.dataCentersRepo.GetDataCenter(ctx, dataCenter)
				if err != nil {
					log.Info("error on getting data center in executing endpoint rule: ", err)
					waitGroup.Done()
					return
				}

				var newHeader = map[string][]string{}
				for key, val := range rule.Header {
					newHeader[key] = strings.Split(val, ",")
				}

				respBody, respHeader, respStatus, respTime, err := e.agentHandler.SendCurl(ctx, dataCenter.Baseurl, usecase_models.AgentCurlRequest{
					Url:    rule.Url,
					Method: rule.Method,
					Header: newHeader,
					Body:   rule.Body,
				})
				if err != nil {
					log.Info("error on sending curl in executing rule: ", err)
					waitGroup.Done()
					return
				}
				if !curlAcceptanceCriteria(respStatus, []byte(respBody), rule.AcceptanceModel) {
					err := e.endpointReportRepo.WriteEndpointReport(ctx, endpointRules.Scheduling.ProjectId, endpointRules.Scheduling.PipelineId, 0, respTime)
					if err != nil {
						log.Info("error on writing curl report in executing rule: ", err)
					}
					// TODO: send alert
					break
				}
				err = e.endpointReportRepo.WriteEndpointReport(ctx, endpointRules.Scheduling.ProjectId, endpointRules.Scheduling.PipelineId, 1, respTime)
				if err != nil {
					log.Info("error on writing curl report in executing rule: ", err)
				}

				responses.BodyResponses[rule.SequenceId] = string(respBody)
				responses.HeaderResponses[rule.SequenceId] = respHeader
			}
			waitGroup.Done()
		}(dataC)
	}

	waitGroup.Wait()
	return nil
}

func curlAcceptanceCriteria(status string, body []byte, acceptRules usecase_models.AcceptanceModel) bool {
	statusCheck := false
	for _, val := range acceptRules.Statuses {
		if val == status {
			statusCheck = true
			break
		}
	}
	if !statusCheck {
		return false
	}

	var respbody map[string]interface{}
	json.Unmarshal(body, &respbody)

	bodyCheck := true
	for _, val := range acceptRules.ResponseBodies {
		_, ok := respbody[val.Key]
		if !ok {
			bodyCheck = false
			break
		}
		//if reflect.TypeOf(value).String() != val.Value {
		//	bodyCheck = false
		//	break
		//}
	}

	if !bodyCheck {
		return false
	}

	return true
}

func getEndpoint(id int, endpointRules []usecase_models.EndpointRules) (usecase_models.EndpointRules, error) {
	for _, rule := range endpointRules {
		if rule.SequenceId == id {
			return rule, nil
		}
	}
	return usecase_models.EndpointRules{}, errors.New("dont know why")
}
