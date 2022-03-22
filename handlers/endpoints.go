package handlers

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"test-manager/models"
)

type EndpointHandler interface {
	RegisterRules(rules models.EndpointRequest) error
}

type endpointHandler struct {
}

func NewEndpointHandler() EndpointHandler {
	return &endpointHandler{}
}

func (e *endpointHandler) RegisterRules(rules models.EndpointRequest) error {
	for _, rule := range rules.Endpoints {
		for _, bodyDependency := range rule.BodyDependency {
			parentEndpoint, err := getEndpoint(bodyDependency.EndpointId, rules.Endpoints)
			if err != nil {
				panic(err)
			}

			value := gjson.Result{}
			if bodyDependency.ParentKey[0:7] == "$header_" {
				value = gjson.Get(parentEndpoint.Body, bodyDependency.ParentKey)
			} else if bodyDependency.ParentKey[0:5] == "$body_" {
				value = gjson.Get(parentEndpoint.Body, bodyDependency.ParentKey)
			} else {
				panic("wtf")
			}
			ruleBody := gjson.Get(rule.Body, "#(...)#").Value().(map[string]interface{})
			ruleBody[bodyDependency.Key] = value.String()
			newBody, err := json.Marshal(ruleBody)
			rule.Body = string(newBody)
		}
		for _, headerDependency := range rule.HeaderDependency {
			parentEndpoint, err := getEndpoint(headerDependency.EndpointId, rules.Endpoints)
			if err != nil {
				panic(err)
			}

			value := gjson.Result{}
			if headerDependency.ParentKey[0:7] == "$header_" {
				value = gjson.Get(parentEndpoint.Body, headerDependency.ParentKey)
			} else if headerDependency.ParentKey[0:5] == "$body_" {
				value = gjson.Get(parentEndpoint.Body, headerDependency.ParentKey)
			} else {
				panic("wtf")
			}
			rule.Header[headerDependency.Key] = value.String()
		}
	}
	return nil
}

func getEndpoint(id int, endpointRules []models.EndpointRules) (models.EndpointRules, error) {
	for _, rule := range endpointRules {
		if rule.ID == id {
			return rule, nil
		}
	}
	return models.EndpointRules{}, errors.New("dont know why")
}
