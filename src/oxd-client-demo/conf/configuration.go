package conf

import "oxd-client/model/params"

type Configuration struct{
	Host string `toml:"host"`
	Key string `toml:"key"`
	Cert string `toml:"cert"`

	RegisterSiteRequestParams model.RegisterSiteRequestParams `toml:"registerSiteRequestParams"`
	UpdateSiteRequestParams model.UpdateSiteRequestParams `toml:"updateSiteRequestParams"`
	AuthorizationCodeRequestParams model.AuthorizationCodeRequestParams `toml:"authorizationCodeRequestParams"`
}
