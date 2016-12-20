package page

import (
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client-demo/service"
	"oxd-client/client"
	"oxd-client/constants"
	"oxd-client/model/params"
	"oxd-client/model/transport"
	"encoding/json"
	"fmt"
)

func UserInfoPage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session conf.SessionVars) {

	var oxdResponse transport.OxdResponse

	page.CallOxdServer(
		client.BuildOxdRequest(constants.GET_USER_INFO,model.UserInfoRequestParams{session.OxdId,session.AccessToken}),
		&oxdResponse,
		configuration.Host)

	var response model.UserInfoResponseParams
	oxdResponse.GetParams(&response)
	display(w,response)
}

func display(w http.ResponseWriter,response model.UserInfoResponseParams){
	value, _ := json.Marshal(response)
	fmt.Fprintln(w,string(value))
}
