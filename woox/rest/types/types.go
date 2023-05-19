package types

type HTTPRequest struct {
	URL     string
	Path    string
	Method  string
	Headers map[string]string
	Query   any
	Body    any
	Debug   bool
}

type Response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"messsage"`
}
