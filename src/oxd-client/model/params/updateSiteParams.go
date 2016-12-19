package model

import "time"

type UpdateSiteRequestParams struct {
	OpHost                        string  `json:"op_host,omitempty"`

	OxdId                         string `json:"oxd_id"`
	AuthorizationRedirectUri      string `json:"authorization_redirect_uri,omitempty"`
	PostLogoutRedirectUri         string  `json:"post_logout_redirect_uri,omitempty"`

	RedirectUris                  []string  `json:"redirect_uris,omitempty"`
	ResponseTypes                 []string  `json:"response_types,omitempty"`

	ClientId                      string  `json:"client_id,omitempty"`
	ClientSecret                  string  `json:"client_secret,omitempty"`
	ClientJwksUri                 string  `json:"client_jwks_uri,omitempty"`
	ClientSectorIdentifierUri     string  `json:"client_sector_identifier_uri,omitempty"`
	ClientTokenEndpointAuthMethod string  `json:"client_token_endpoint_auth_method,omitempty"`
	ClientRequestUris             []string  `json:"client_request_uris,omitempty"`
	ClientLogoutUri               []string  `json:"client_logout_uris,omitempty"`
	ClientSecretExpiresAt         time.Time  `json:"client_secret_expires_at,omitempty"`

	Scope                         []string  `json:"scope,omitempty"`
	UiLocales                     []string  `json:"ui_locales,omitempty"`
	ClaimsLocales                 []string  `json:"claims_locales,omitempty"`
	AcrValues                     []string  `json:"acr_values,omitempty"`
	GrantType                     []string  `json:"grant_types,omitempty"`
	Contacts                      []string  `json:"contacts,omitempty"`
}

type UpdateSiteResponseParams struct {
	OxdId string `json:"oxd_id"`
}