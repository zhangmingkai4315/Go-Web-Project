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
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
