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

func main() {

	loggo.GetLogger("transport.transport").SetLogLevel(loggo.DEBUG)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		page.RegisterSitePage(w, r, serverConf)
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		page.UpdateSitePage(w, r, serverConf)
	})

	http.HandleFunc("/authUrl", func(w http.ResponseWriter, r *http.Request) {
		page.AuthorizationUrlPageSite(w, r, serverConf)
	})

	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}

func init() {
	_, err := toml.DecodeFile("conf/main.toml", &serverConf)
	utils.CheckError("transport.transport","Config file error",err)
}
