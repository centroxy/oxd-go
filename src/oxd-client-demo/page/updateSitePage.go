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


func UpdateSitePage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration) {

	result := client.BuildOxdResponse(model.UpdateSiteResponseParams{})
	client.Send(
		client.BuildOxdRequest(constants.UPDATE_SITE,configuration.UpdateSiteRequestParams),
		configuration.Host, &result)

	response,err := json.Marshal(result)
	utils.CheckError("UpdateSitePage","Marshal error",err)
	value, err := json.Marshal(configuration.RegisterSiteRequestParams)
	utils.CheckError("UpdateSitePage","Marshal error",err)
	fmt.Fprintf(w, string(value))
	fmt.Fprintf(w, string(response))
}
