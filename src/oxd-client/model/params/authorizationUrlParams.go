package model

type AuthorizationUrlRequestParams struct {
	OxdId string `json:"oxd_id"`
	AcrValues []string  `json:"acr_values"`
	Prompt string `json:"prompt"`
	Scope []string `json:"scope,omitempty"`
}

type AuthorizationUrlResponseParams struct {
	AuthorizationUrl string `json:"authorization_url"`
}
