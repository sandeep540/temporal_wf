package model

type Interceptor []struct {
	WorkflowId string `json:"workflowId"`
	Activity   string `json:"activity"`
	ServiceURL string `json:"serviceURL"`
}
