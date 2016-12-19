package model

type AuthorizationCodeRequestParams struct {
	OxdId string `json:"oxd_id"`
	RedirectUrl string `json:"redirect_url,omitempty"`
	ClientId string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	UserId string `json:"user_id,omitempty"`
	UserSecret string `json:"user_secret,omitempty"`
	Scope []string `json:"scope,omitempty"`
	Nonce string `json:"nonce,omitempty"`
	Acr string `json:"acr,omitempty"`
}

type AuthorizationCodeResponseParams struct {
	AuthorizationUrl string `json:"authorization_url"`
}

