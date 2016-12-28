package page

import (
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params/url"
	"oxd-client/constants"
	"oxd-client/model/transport"
	"oxd-client-demo/service"
	"github.com/juju/loggo"
)

func AuthorizationUrlPageSite(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session conf.SessionVars) {

	log :=loggo.GetLogger("default")
	var oxdResponse transport.OxdResponse

	page.CallOxdServer(
		client.BuildOxdRequest(constants.GET_AUTHORIZATION_URL,
			model.AuthorizationUrlRequestParams{session.OxdId,make([]string, 0),"",make([]string, 0)}),
		&oxdResponse,
		configuration.Host)


	var response model.AuthorizationUrlResponseParams
	oxdResponse.GetParams(&response)

	log.Debugf(response.AuthorizationUrl)
	http.Redirect(w, r, response.AuthorizationUrl, http.StatusPermanentRedirect)
}
