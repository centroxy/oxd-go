package tests

import (
	"testing"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"
	"oxd-client/model/params"
	"oxd-client-test/utils"
	"oxd-client/model/transport"
	"oxd-client-test/conf"
)

func TestGetUserInfo(t *testing.T) {
	//BEFORE
	codeResponse, oxdId := utils.ExecCodeFlow()
	requestParams := model.UserInfoRequestParams{oxdId,codeResponse.AccessToken}
	request := client.BuildOxdRequest(constants.GET_USER_INFO,requestParams)
	var response transport.OxdResponse
	var responseParams model.UserInfoResponseParams

	//TEST
	client.Send(request,conf.TestConfiguration.Host,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.Claims,"AccessToken should not be empty")
}