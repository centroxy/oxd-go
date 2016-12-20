package main

import (
	"oxd-client-demo/page"
	"net/http"
	"log"
	"github.com/juju/loggo"
	"github.com/BurntSushi/toml"
	"oxd-client/utils"
	"oxd-client-demo/conf"
)


var serverConf conf.Configuration
var session conf.SessionVars

func main() {

	loggo.GetLogger("default").SetLogLevel(loggo.DEBUG)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		page.RegisterSitePage(w, r, serverConf, &session)
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		page.UpdateSitePage(w, r, serverConf, &session)
	})

	http.HandleFunc("/authUrl", func(w http.ResponseWriter, r *http.Request) {
		page.AuthorizationUrlPageSite(w, r, serverConf,session)
	})

	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		page.RedirectPage(w, r, serverConf, &session)
	})

	http.HandleFunc("/userInfo", func(w http.ResponseWriter, r *http.Request) {
		page.UserInfoPage(w, r, serverConf, session)
	})

	http.HandleFunc("/authCode", func(w http.ResponseWriter, r *http.Request) {
		page.AuthorizationCodePageSite(w, r, serverConf)
	})

	err := http.ListenAndServeTLS(":8080",serverConf.Cert, serverConf.Key, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}

func init() {
	_, err := toml.DecodeFile("conf/main.toml", &serverConf)
	utils.CheckError("transport.transport","Config file error",err)
}
