package tests

import (
	"testing"
	"oxd-client/model/transport"
	"oxd-client/model/params/registration"
	"oxd-client-test/conf"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"

	"oxd-client-test/utils"
)

func TestRegisterClientSite(t *testing.T) {
	//BEFORE
	requestParams := utils.PrepareRegisterParams()
	request := client.BuildOxdRequest(constants.REGISTER_SITE,requestParams)
	var response transport.OxdResponse
	var responseParams model.RegisterSiteResponseParams

	//TEST
	client.Send(request,conf.TestConfiguration.Host,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.OpHost,"Host should not be empty")
	assert.NotEmpty(t,responseParams.OxdId,"OxdId should not be empty")
}

func TestUpdateClientSite(t *testing.T) {
	//BEFORE
	updateParams := utils.PrepareUpdateParams(utils.RegisterClient())
	updateRequest := client.BuildOxdRequest(constants.UPDATE_SITE,updateParams)
	var updateResponse transport.OxdResponse
	var updateResponseParams model.UpdateSiteResponseParams

	//TEST
	client.Send(updateRequest,conf.TestConfiguration.Host,&updateResponse)

	//ASSERT
	updateResponse.GetParams(&updateResponseParams)
	assert.Equal(t,constants.STATUS_OK,updateResponse.Status,"Status should be ok")
	assert.NotEmpty(t,updateResponseParams.OxdId,"OxdId should not be empty")
}

