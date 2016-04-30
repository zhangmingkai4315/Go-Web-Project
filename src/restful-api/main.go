package main

import (
	"github.com/codegangsta/negroni"
	"log"
	"net/http"
	"restful-api/common"
	"restful-api/routers"
)

func main() {
	common.StartUp()
	router := routers.InitRouters()
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr:   common.AppConfig.Server,
		Handle: n,
	}
	log.Prinln("Listening...")
	server.ListenAndServe()
}
