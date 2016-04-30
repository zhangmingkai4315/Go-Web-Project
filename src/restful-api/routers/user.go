package routers

import (
	"github.com/gotilla/mux"
	"restful-api/controllers"
)

func SetUserRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/user/register", controllers.Register).Methods("POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
}
