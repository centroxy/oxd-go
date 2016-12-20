package page

import (
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client-demo/service"
	"oxd-client/client"
	"oxd-client/constants"
	"oxd-client/model/params"
	"oxd-client/model/transport"
)


func UpdateSitePage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars) {
	var oxdResponse transport.OxdResponse

	page.CallOxdServer(
		client.BuildOxdRequest(constants.UPDATE_SITE,configuration.UpdateSiteRequestParams),
		&oxdResponse,
		configuration.Host)

	var params model.UpdateSiteRequestParams
	oxdResponse.GetParams(&params)
	session.OxdId = params.OxdId
}
