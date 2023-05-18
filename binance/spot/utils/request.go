package utils

type HTTPRequest struct {
	// SecurityType each endpoint has a security type that determines how you will interact with it
	// docs: https://binance-docs.github.io/apidocs/spot/en/#endpoint-security-type
	SecurityType SecurityType

	BaseURL string
	Path    string
	Method  string
	Headers map[string]string
	Query   interface{}
	Body    interface{}
}
