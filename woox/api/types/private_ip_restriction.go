package types

type IPRestriction struct {
	Response
	Rows []struct {
		IPList     string `json:"ip_list"`
		APIKey     string `json:"api_key"`
		UpdateTime string `json:"update_time"`
		Restrict   bool   `json:"restrict"`
	} `json:"rows"`
	Meta struct {
		Total          int `json:"total"`
		RecordsPerPage int `json:"records_per_page"`
		CurrentPage    int `json:"current_page"`
	} `json:"meta"`
}
