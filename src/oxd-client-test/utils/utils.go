//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package utils

import (
	"oxd-client/client"
	"oxd-client/constants"

	"oxd-client-test/conf"
	"oxd-client/model/transport"
	url "oxd-client/model/params/url"
	token "oxd-client/model/params/token"
	"oxd-client/model/params/registration"
)

func RegisterClient() string{
	requestParams := PrepareRegisterParams()
	request := client.BuildOxdRequest(constants.REGISTER_SITE,requestParams)
	var response transport.OxdResponse
	var responseParams model.RegisterSiteResponseParams

	client.Send(request,conf.TestConfiguration.Host,&response)

	response.GetParams(&responseParams)
	return responseParams.OxdId
}

func PrepareRegisterParams()  model.RegisterSiteRequestParams {
	var requestParams model.RegisterSiteRequestParams
	requestParams.OpHost = conf.TestConfiguration.OpHost
	requestParams.AuthorizationRedirectUri = conf.TestConfiguration.RedirectUrl
	requestParams.PostLogoutRedirectUri = conf.TestConfiguration.PostLogoutRedirectUrl
	requestParams.ClientLogoutUri = []string {conf.TestConfiguration.LogoutUrl}
	requestParams.RedirectUris = [] string {conf.TestConfiguration.RedirectUrl}
	requestParams.AcrValues = make([]string,0)
	requestParams.Scope = []string{"openid", "profile"}
	requestParams.GrantType = []string{"authorization_code"}
	requestParams.ResponseTypes = []string{"code"}
	return  requestParams
}

func PrepareUpdateParams(oxd_id string)  model.UpdateSiteRequestParams {
	var requestParams model.UpdateSiteRequestParams
	requestParams.OxdId = oxd_id
	requestParams.OpHost = conf.TestConfiguration.OpHost
	requestParams.AuthorizationRedirectUri = conf.TestConfiguration.RedirectUrl
	requestParams.PostLogoutRedirectUri = conf.TestConfiguration.PostLogoutRedirectUrl
	requestParams.ClientLogoutUri = []string {conf.TestConfiguration.LogoutUrl}
	requestParams.RedirectUris = [] string {conf.TestConfiguration.RedirectUrl}
	requestParams.AcrValues = make([]string,0)
	requestParams.Scope = []string{"openid"}
	requestParams.GrantType = []string{"authorization_code"}
	requestParams.ResponseTypes = []string{"code"}
	return  requestParams
}

func ExecCodeFlow() (url.AuthorizationCodeFlowResponseParams,string){
	requestParams := url.AuthorizationCodeFlowRequestParams{
		RegisterClient(),
		conf.TestConfiguration.RedirectUrl,
		conf.TestConfiguration.ClientId,
		conf.TestConfiguration.ClientSecret,
		conf.TestConfiguration.UserId,
		conf.TestConfiguration.UserSecret,
		conf.TestConfiguration.Scope,"",""}
	request := client.BuildOxdRequest(constants.AUTHORIZATION_CODE_FLOW,requestParams)
	var response transport.OxdResponse
	var responseParams url.AuthorizationCodeFlowResponseParams
	client.Send(request,conf.TestConfiguration.Host,&response)
	response.GetParams(&responseParams)
	return responseParams, requestParams.OxdId
}


func ExecCode() (token.AuthorizationCodeResponseParams, token.AuthorizationCodeRequestParams) {
	requestParams := token.AuthorizationCodeRequestParams{RegisterClient(),
		conf.TestConfiguration.UserId,
		conf.TestConfiguration.UserSecret,
		make([]string,0),
		"","af0ifjsldkj"}
	request := client.BuildOxdRequest(constants.GET_AUTHORIZATION_CODE,requestParams)
	var response transport.OxdResponse
	var responseParams token.AuthorizationCodeResponseParams
	client.Send(request,conf.TestConfiguration.Host,&response)
	response.GetParams(&responseParams)
	return responseParams, requestParams
}