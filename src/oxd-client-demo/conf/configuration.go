package conf

import "oxd-client/model/params"

type Configuration struct{
	Host string `toml:"host"`
	RegisterSiteRequestParams model.RegisterSiteRequestParams `toml:"registerSiteRequestParams"`
	UpdateSiteRequestParams model.UpdateSiteRequestParams `toml:"updateSiteRequestParams"`
	AuthorizationUrlRequestParams model.AuthorizationUrlRequestParams `toml:"authorizationUrlRequestParams"`
}
