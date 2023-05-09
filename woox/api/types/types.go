package types

type Response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"messsage"`
}
