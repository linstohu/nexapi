package types

type SubAccounts struct {
	Response
	Rows []struct {
		ApplicationID string `json:"application_id"`
		Account       string `json:"account"`
		CreatedTime   string `json:"created_time"`
	} `json:"rows"`
}
