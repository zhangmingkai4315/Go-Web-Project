package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandle struct {
	message string
}

func (m *messageHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()
	mh1 := &messageHandle{
		"Welcome to go web develop",
	}
	mux.Handle("/welcome", mh1)

	mh2 := &messageHandle{
		"Hello this is mike",
	}
	mux.Handle("/hello", mh2)

	log.Println("Listening ...")
	http.ListenAndServe(":8080", mux)
}
