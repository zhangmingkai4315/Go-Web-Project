package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"restful-api/common"
	"restful-api/controllers"
)

func SetTaskRouter(router *mux.Router) *mux.Router {
	taskRouter := mux.NewRouter()
	taskRouter.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")

	taskRouter.HandleFunc("/tasks/{id}", controllers.UpdateTaskById).Methods("PUT")
	taskRouter.HandleFunc("/tasks/{id}", controllers.GetTaskById).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}", controllers.DeleteTaskById).Methods("DELETE")

	taskRouter.HandleFunc("/tasks/users/{id}", controllers.GetTasksByUser).Methods("GET")

	router.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandleFunc(common.Authorize),
		negroni.Wrap(taskRouter),
	))
	return router
}
