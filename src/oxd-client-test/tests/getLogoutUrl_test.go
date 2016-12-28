package tests

import (
	"testing"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"
	"oxd-client/model/params/url"
	"oxd-client/model/transport"
	"oxd-client-test/conf"
	"oxd-client-test/utils"
)

func TestGetLogoutUrl(t *testing.T) {
	//BEFORE
	codeResponse, oxdId := utils.ExecCodeFlow()
	requestParams := model.LogoutUrlRequestParams{oxdId, codeResponse.IdToken,
		conf.TestConfiguration.PostLogoutRedirectUrl, "",""}
	request := client.BuildOxdRequest(constants.GET_LOGOUT_URI,requestParams)
	var response transport.OxdResponse
	var responseParams model.LogoutUrlResponseParams

	//TEST
	client.Send(request,conf.TestConfiguration.Host,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.Uri,"Uri should not be empty")
}
