package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}
func message(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "this is a message")
}
func middlewareFirst(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	log.Println("middlewareFirst - before")
	context.Set(req, "user", "mike")
	next(w, req)
	log.Println("middlewareFirst - after")
}

func middlewareSecond(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	log.Println("middlewareSecond - before")
	fmt.Printf("The user is :%+v", context.Get(r, "user"))
	if r.URL.Path == "/message" {
		if r.URL.Query().Get("password") == "password" {
			log.Println("Authorized to the system")
			next(w, r)
		} else {
			log.Println("Failed")
			return
		}
	} else {
		next(w, r)
	}

	// next(w, req)
	log.Println("middlewareSecond - after")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/message", message)
	n := negroni.Classic()
	// n.UseHandler(mux)
	n.Use(negroni.HandlerFunc(middlewareFirst))
	n.Use(negroni.HandlerFunc(middlewareSecond))
	n.UseHandler(mux)
	n.Run(":8080")
}
