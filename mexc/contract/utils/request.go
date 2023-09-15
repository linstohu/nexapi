package ctutils

type HTTPRequest struct {
	BaseURL string
	Path    string
	Method  string
	Headers map[string]string
	Query   any
	Body    any
}
