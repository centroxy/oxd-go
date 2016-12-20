package page

import (
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params"
	"oxd-client/constants"
	"oxd-client-demo/service"
	"oxd-client/model/transport"
	"encoding/json"
	"fmt"
)

func AuthorizationCodePageSite(w http.ResponseWriter, r *http.Request, configuration conf.Configuration) {
	var oxdResponse transport.OxdResponse

	page.CallOxdServer(
		client.BuildOxdRequest(constants.AUTHORIZATION_CODE_FLOW,configuration.AuthorizationCodeRequestParams),
		&oxdResponse,
		configuration.Host)

	var response model.AuthorizationCodeResponseParams
	oxdResponse.GetParams(&response)
	displayCodeResponseParams(w,response)
}

func displayCodeResponseParams(w http.ResponseWriter,response model.AuthorizationCodeResponseParams){
	value, _ := json.Marshal(response)
	fmt.Fprintln(w,string(value))
}

