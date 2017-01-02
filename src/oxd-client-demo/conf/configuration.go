//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package conf

import "oxd-client/model/params/registration"
import param "oxd-client/model/params/url"


type Configuration struct{
	Host string `toml:"host"`
	Key string `toml:"key"`
	Cert string `toml:"cert"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Loglevel string `toml:"loglevel"`


	RegisterSiteRequestParams model.RegisterSiteRequestParams `toml:"registerSiteRequestParams"`
	UpdateSiteRequestParams model.UpdateSiteRequestParams `toml:"updateSiteRequestParams"`
	AuthorizationCodeFlowRequestParams param.AuthorizationCodeFlowRequestParams `toml:"authorizationCodeFlowRequestParams"`
}
