package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kr/pretty"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
	"test-manager/models"
)

type EndpointHandler interface {
	RegisterRules(ctx context.Context, rules models.EndpointRequest) error
}

type endpointHandler struct {
}

func NewEndpointHandler() EndpointHandler {
	return &endpointHandler{}
}

func (e *endpointHandler) RegisterRules(ctx context.Context, rules models.EndpointRequest) error {
	fmt.Printf("%# v", pretty.Formatter(rules))
	var responses = models.EndpointResponses{}
	for _, rule := range rules.Endpoints {
		var value []string
		for _, bodyDependency := range rule.BodyDependency {
			if bodyDependency.ParentKey[0:8] == "$header_" {
				value = responses.HeaderResponses[bodyDependency.EndpointId][bodyDependency.ParentKey[8:]]
			} else if bodyDependency.ParentKey[0:6] == "$body_" {
				value = []string{gjson.Get(responses.BodyResponses[bodyDependency.EndpointId], bodyDependency.ParentKey[6:]).String()}
			} else {
				panic("wtf")
			}
			ruleBody := gjson.Get(rule.Body, "#(...)#").Value().(map[string]interface{})
			ruleBody[bodyDependency.Key] = strings.Join(value[:], ",")
			newBody, err := json.Marshal(ruleBody)
			if err != nil {
				panic(err)
			}
			rule.Body = string(newBody)
		}
		for _, headerDependency := range rule.HeaderDependency {
			var value []string
			if headerDependency.ParentKey[0:8] == "$header_" {
				value = responses.HeaderResponses[headerDependency.EndpointId][headerDependency.ParentKey[8:]]
			} else if headerDependency.ParentKey[0:6] == "$body_" {
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
		responses.BodyResponses[rule.ID] = string(respBody)
		responses.HeaderResponses[rule.ID] = resp.Header
	}
	fmt.Printf("%# v", pretty.Formatter(rules))
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
