package types

type Response struct {
	Success bool `json:"success"`
	Code    int  `json:"code"`
}

type ServerTime struct {
	Response
	Data int64 `json:"data"`
}
