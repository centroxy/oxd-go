//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package page

import (
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client-demo/service"
	"oxd-client/client"
	"oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client/model/params/registration"
	"oxd-client-demo/utils"
)


func UpdateSitePage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars) {
	var oxdResponse transport.OxdResponse

	page.CallOxdServer(
		client.BuildOxdRequest(constants.UPDATE_SITE,configuration.UpdateSiteRequestParams),
		&oxdResponse,
		configuration.Host)

	var response model.UpdateSiteResponseParams
	oxdResponse.GetParams(&response)
	session.OxdId = response.OxdId
	utils.DisplayResponse(w,response)
}
