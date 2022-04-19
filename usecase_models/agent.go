package usecase_models

type AgentCurlRequest struct {
	Url    string              `json:"url"`
	Method string              `json:"method"`
	Header map[string][]string `json:"header"`
	Body   string              `json:"body"`
}
