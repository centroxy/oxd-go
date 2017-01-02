//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package model

type RegisterSiteRequestParams struct {

	OpHost string  `json:"op_host,omitempty"`
	AuthorizationRedirectUri string  `json:"authorization_redirect_uri"`
	PostLogoutRedirectUri string  `json:"post_logout_redirect_uri,omitempty"`

	RedirectUris []string  `json:"redirect_uris,omitempty"`
	ResponseTypes []string  `json:"response_types,omitempty"`

	ClientId string  `json:"client_id,omitempty"`
	ClientSecret string  `json:"client_secret,omitempty"`
	ClientName string  `json:"client_name,omitempty"`
	ClientJwksUri string  `json:"client_jwks_uri,omitempty"`
	ClientTokenEndpointAuthMethod string  `json:"client_token_endpoint_auth_method,omitempty"`
	ClientRequestUris  []string  `json:"client_request_uris,omitempty"`
	ClientLogoutUri  []string  `json:"client_logout_uris,omitempty"`
	ClientSectorIdentifierUri string  `json:"client_sector_identifier_uri,omitempty"`

	Scope  []string  `json:"scope,omitempty"`
	UiLocales  []string  `json:"ui_locales,omitempty"`
	ClaimsLocales  []string  `json:"claims_locales,omitempty"`
	AcrValues  []string  `json:"acr_values,omitempty"`
	GrantType  []string  `json:"grant_types,omitempty"`
	Contacts []string  `json:"contacts,omitempty"`
}

type RegisterSiteResponseParams struct {
	OxdId string `json:"oxd_id"`
	OpHost string `json:"op_host"`
}