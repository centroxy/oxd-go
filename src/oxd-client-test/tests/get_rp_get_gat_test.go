//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package tests

import (
	"testing"
	"oxd-client/client"
	"oxd-client/constants"
	"github.com/stretchr/testify/assert"
	"oxd-client-test/utils"
	"oxd-client/model/params/uma"
	"oxd-client/model/transport"
	"oxd-client-test/conf"
)

func TestRpGetGat(t *testing.T) {
	//BEFORE
	requestParams := uma.RpGetGatRequestParams{utils.RegisterClient(),[]string{"http://photoz.example.com/dev/actions/all"}}
	request := client.BuildOxdRequest(constants.RP_GET_GAT,requestParams)
	var response transport.OxdResponse
	var responseParams uma.RpGetGatResponseParams

	//TEST
	client.Send(request,conf.TestConfiguration.Host,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.Rpt,"Rpt should not be empty")
}