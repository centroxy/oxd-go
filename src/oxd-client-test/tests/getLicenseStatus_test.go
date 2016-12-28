package tests

import (
	"testing"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"
	"oxd-client/model/params/validation"
	"oxd-client/model/transport"
	"oxd-client-test/conf"
)

func TestLicenseStatus(t *testing.T) {
	//BEFORE
	requestParams := validation.LicenseStatusRequestParams{}
	request := client.BuildOxdRequest(constants.LICENSE_STATUS,requestParams)
	var response transport.OxdResponse
	var responseParams validation.LicenseStatusResponseParams

	//TEST
	client.Send(request,conf.TestConfiguration.Host,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.Equal(t,true,responseParams.Valid,"License not valid")
	assert.NotEmpty(t,responseParams.Features,"AuthorizationUrl should not be empty")
	assert.NotEmpty(t,responseParams.Name,"Name should not be empty")
	assert.NotEmpty(t,responseParams.ThreadCount,"ThreadCount should not be empty")
}
