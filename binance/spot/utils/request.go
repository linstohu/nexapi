package utils

type HTTPRequest struct {
	BaseURL string
	Path    string
	Method  string
	Headers map[string]string
	Query   interface{}
	Body    interface{}
}
