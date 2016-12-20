package page

import (
	"net/http"
	"oxd-client-demo/conf"
	"oxd-client/client"
	"oxd-client/model/params"
	"oxd-client/constants"
	"oxd-client/model/transport"
	"encoding/json"
	"fmt"
	"oxd-client-demo/service"
)

func RedirectPage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars) {

	readQueryParams(r,session)
	response := getTokensByCode(*session,configuration)
	saveTokens(response,session)
	displayResponse(w,response)
}

func readQueryParams(r *http.Request, session *conf.SessionVars){
	session.State = r.URL.Query().Get("state")
	session.Code = r.URL.Query().Get("code")
}

func getTokensByCode(session conf.SessionVars,configuration conf.Configuration) model.TokensByCodeResponseParams{
	var oxdResponse transport.OxdResponse
	page.CallOxdServer(
		client.BuildOxdRequest(constants.GET_TOKENS_BY_CODE,
			model.TokensByCodeRequestParams{session.OxdId, session.Code, session.State}),
		&oxdResponse,
		configuration.Host)

	var response model.TokensByCodeResponseParams
	oxdResponse.GetParams(&response)
	return response
}


func saveTokens(response model.TokensByCodeResponseParams, session *conf.SessionVars){
	session.AccessToken = response.AccessToken
	session.IdToken = response.IdToken
}

func displayResponse(w http.ResponseWriter,response model.TokensByCodeResponseParams){
	value, _ := json.Marshal(response)
	fmt.Fprintln(w,string(value))
}