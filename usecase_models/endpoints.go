package usecase_models

type EndpointRequest struct {
	Endpoints  []EndpointRules `json:"endpoints"`
	Scheduling Scheduling      `json:"scheduling"`
}

type EndpointRules struct {
	ID               int               `json:"id"`
	Url              string            `json:"url"`
	Method           string            `json:"method"`
	Body             string            `json:"body"`
	BodyDependency   []Dependency      `json:"body_dependency"`
	Header           map[string]string `json:"header"`
	HeaderDependency []Dependency      `json:"header_dependency"`
	AcceptanceModel  AcceptanceModel   `json:"acceptance_model"` // check keys with their type and status
}

type Scheduling struct {
	ProjectId      int    `json:"project_id"`
	Duration       int    `json:"duration"`
	EndAt          string `json:"end_at"`
	DataCentersIds []int  `json:"data_centers"` // datacenter id
}

type AcceptanceModel struct {
	Statuses       []string        `json:"statuses"`
	ResponseBodies []KeyValueModel `json:"response_bodies"`
}

type KeyValueModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Dependency struct {
	EndpointId int `json:"endpoint_id"`

	// parent key and key works with gjson format
	// like key1.key2.key3
	// this gives us key3 that is inside key2 and that is inside key1
	ParentKey string `json:"parent_key"` // $header_... for headers $body_... for bodies
	Key       string `json:"key"`

	//KeyParentType string `json:"key_parent_type"`
	//KeyType       string `json:"key_type"`
}

type EndpointResponses struct {
	HeaderResponses map[int]map[string][]string `json:"header_responses"`
	BodyResponses   map[int]string              `json:"body_responses"`
}
