package page

import (
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params"
	"oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client-demo/service"
)

func RegisterSitePage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars) {
	var oxdResponse transport.OxdResponse

	page.CallOxdServer(
		client.BuildOxdRequest(constants.REGISTER_SITE,configuration.RegisterSiteRequestParams),
		&oxdResponse,
		configuration.Host)

	var params model.RegisterSiteResponseParams
	oxdResponse.GetParams(&params)
	session.OxdId = params.OxdId
}



