package models

type EndpointRequest struct {
	Endpoints []EndpointRules `json:"endpoints"`
}

type EndpointRules struct {
	ID               int          `json:"id"`
	Url              string       `json:"url"`
	Body             interface{}  `json:"body"`
	BodyDependency   []Dependency `json:"body_dependency"`
	Header           interface{}  `json:"header"`
	HeaderDependency []Dependency `json:"header_dependency"`
}

type Dependency struct {
	EndpointId    int    `json:"endpoint_id"`
	ParentKey     string `json:"parent_key"`
	Key           string `json:"key"` // $header_... for headers $body_... for bodies
	KeyParentType string `json:"key_parent_type"`
	KeyType       string `json:"key_type"`
}
