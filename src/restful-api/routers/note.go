package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"restful-api/common"
	"restful-api/controllers"
)

func SetNoteRoutes(router *mux.Router) *mux.Router {
	noteRouter := mux.NewRouter()
	noteRouter.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
	noteRouter.HandleFunc("/notes/{id}", controllers.UpdateNote).Methods("PUT")
	noteRouter.HandleFunc("/notes/{id}", controllers.GetNoteById).Methods("GET")
	noteRouter.HandleFunc("/notes/{id}", controllers.DeleteNoteById).Methods("DELETE")

	noteRouter.HandleFunc("/notes/tasks/{id}", controllers.GetNotesByTask).Methods("GET")
	noteRouter.HandleFunc("/notes", controllers.GetNotes).Methods("GET")

	router.PathPrefix("/notes").Handler(negroni.New(
		negroni.HandleFunc(common.Authorize),
		negroni.Wrap(noteRouter),
	))
	return router
}
