package page

import (
	"net/http"
	"fmt"
	"oxd-client/model/params"
	"encoding/json"
	"oxd-client/utils"
	"oxd-client/client"
	"oxd-client/constants"
	"oxd-client-demo/conf"
)


func RegisterSitePage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration) {

	result := client.BuildOxdResponse(model.RegisterSiteResponseParams{})
	client.Send(
		client.BuildOxdRequest(constants.REGISTER_SITE,configuration.RegisterSiteRequestParams),
		configuration.Host, &result)

	response,err := json.Marshal(result)
	utils.CheckError("RegisterSitePage","Marshal error",err)
	value, err := json.Marshal(configuration.RegisterSiteRequestParams)
	utils.CheckError("RegisterSitePage","Marshal error",err)
	fmt.Fprintf(w, string(value))
	fmt.Fprintf(w, string(response))
}

