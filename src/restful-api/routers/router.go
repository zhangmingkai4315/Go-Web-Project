package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetUserRouters(router)
	router = SetTaskRouter(router)
	router = SetNoteRoutes(router)
	return router
}
