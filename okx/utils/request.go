package okxutils

type HTTPRequest struct {
	BaseURL string
	Path    string
	Method  string
	Headers map[string]string
	Query   any
	Body    any
}

type Response struct {
	Code    string `json:"code"`
	Message string `json:"messsage"`
}
