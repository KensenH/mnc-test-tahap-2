package http

type Response struct {
	Status  string      `json:"status,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty"`
}
