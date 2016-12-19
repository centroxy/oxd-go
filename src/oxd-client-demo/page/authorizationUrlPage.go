package page

import (
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params"
	"oxd-client/constants"
	"encoding/json"
	"oxd-client/utils"
	"fmt"
)

func AuthorizationUrlPageSite(w http.ResponseWriter, r *http.Request, configuration conf.Configuration) {

	result := client.BuildOxdResponse(model.AuthorizationUrlResponseParams{})
	client.Send(
		client.BuildOxdRequest(constants.GET_AUTHORIZATION_URL,configuration.AuthorizationUrlRequestParams),
		configuration.Host, &result)

	response,err := json.Marshal(result)
	utils.CheckError("AuthorizationPageSite","Marshal error",err)
	value, err := json.Marshal(configuration.AuthorizationUrlRequestParams)
	utils.CheckError("AuthorizationPageSite","Marshal error",err)
	fmt.Fprintf(w, string(value))
	fmt.Fprintf(w, string(response))
}
