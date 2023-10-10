package auth

type AuthParams struct {
	GrantType    string `json:"grant_type" validate:"required"`
	ClientID     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	State        string `json:"state"`
	TokenType    string `json:"token_type"`
}
