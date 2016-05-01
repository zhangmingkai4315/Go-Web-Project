package routers

import (
	"github.com/gorilla/mux"
	"log"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetUserRouters(router)
	log.Println("[Routers:Router:SetUserRouters] success----->")
	router = SetTaskRouters(router)
	log.Println("[Routers:Router:SetTaskRouters] success----->")
	router = SetNoteRoutes(router)
	log.Println("[Routers:Router:SetNoteRoutes] success----->")
	return router
}
