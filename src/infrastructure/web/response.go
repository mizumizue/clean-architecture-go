package web

type HttpResponse struct {
	HttpCode int         `json:"code"`
	Response interface{} `json:"response,omitempty"`
	Errors   []string    `json:"errors,omitempty"`
}
