package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/volatiletech/null/v8"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"test-manager/repos"
	"test-manager/tasks/push"
	"test-manager/usecase_models"
	models "test-manager/usecase_models/boiler"
)

type EndpointHandler interface {
	RegisterRules(ctx context.Context, rules usecase_models.EndpointRequest, projectId int) error
	ExecuteRule(ctx context.Context, rules usecase_models.EndpointRequest) error
}

type endpointHandler struct {
	endpointRepo repos.EndpointRepository
	taskPusher   push.TaskPusher
}

func NewEndpointHandler(endpointRepo repos.EndpointRepository, taskPusher push.TaskPusher) EndpointHandler {
	return &endpointHandler{endpointRepo: endpointRepo, taskPusher: taskPusher}
}

func (e endpointHandler) RegisterRules(ctx context.Context, rules usecase_models.EndpointRequest, projectId int) error {
	j, _ := json.Marshal(rules)
	rulesStr := string(j)
	_, err := e.endpointRepo.SaveEndpoint(ctx, models.Endpoint{
		Data:      null.NewString(rulesStr, true),
		ProjectID: projectId,
	})
	if err != nil {
		return err
	}

	_, err = e.taskPusher.PushToEndpoint(ctx, rules)
	if err != nil {
		return err
	}
	return nil
}

func (e *endpointHandler) ExecuteRule(ctx context.Context, rules usecase_models.EndpointRequest) error {
	var responses = usecase_models.EndpointResponses{
		HeaderResponses: map[int]map[string][]string{},
		BodyResponses:   map[int]string{},
	}
	for _, rule := range rules.Endpoints {
		var value []string
		for _, bodyDependency := range rule.BodyDependency {
			if len(bodyDependency.ParentKey) >= 8 && bodyDependency.ParentKey[0:8] == "$header_" {
				value = responses.HeaderResponses[bodyDependency.EndpointId][bodyDependency.ParentKey[8:]]
			} else if len(bodyDependency.ParentKey) >= 6 && bodyDependency.ParentKey[0:6] == "$body_" {
				value = []string{gjson.Get(responses.BodyResponses[bodyDependency.EndpointId], bodyDependency.ParentKey[6:]).String()}
			} else {
				panic("wtf")
			}
			ruleBody := gjson.Parse(rule.Body).Value().(map[string]interface{})
			ruleBody[bodyDependency.Key] = strings.Join(value[:], ",")
			newBody, err := json.Marshal(ruleBody)
			if err != nil {
				panic(err)
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
				panic("wtf")
			}
			rule.Header[headerDependency.Key] = strings.Join(value[:], ",")
		}
		req, err := http.NewRequestWithContext(ctx, rule.Method, rule.Url, bytes.NewBuffer([]byte(rule.Body)))
		if err != nil {
			panic(err)
		}
		for key, value := range rule.Header {
			req.Header.Set(key, value)
		}

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		respBody, _ := ioutil.ReadAll(resp.Body)
		if !acceptanceCriteria(resp.Status, respBody, rule.AcceptanceModel) {
			// TODO: send alert
			break
		}

		responses.BodyResponses[rule.ID] = string(respBody)
		responses.HeaderResponses[rule.ID] = resp.Header
	}
	return nil
}

func acceptanceCriteria(status string, body []byte, acceptRules usecase_models.AcceptanceModel) bool {
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
		value, ok := respbody[val.Key]
		if !ok {
			bodyCheck = false
			break
		}
		if reflect.TypeOf(value).String() != val.Value {
			bodyCheck = false
			break
		}
	}

	if !bodyCheck {
		return false
	}

	return true
}

func getEndpoint(id int, endpointRules []usecase_models.EndpointRules) (usecase_models.EndpointRules, error) {
	for _, rule := range endpointRules {
		if rule.ID == id {
			return rule, nil
		}
	}
	return usecase_models.EndpointRules{}, errors.New("dont know why")
}
