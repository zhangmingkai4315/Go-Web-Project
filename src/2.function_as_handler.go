package main

import (
	"fmt"
	"log"
	"net/http"
)

// func messageHandle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to go web development")
// }

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}

func userHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User:Mike")
}
func main() {
	mux := http.NewServeMux()
	// mh := http.HandlerFunc(messageHandle)
	// mux.Handle("/welcome", mh)

	mux.Handle("/hello", messageHandler("Hello world"))
	mux.Handle("/message", messageHandler("This is a message hander"))

	mux.HandleFunc("/user", userHander)
	log.Println("Listening....")
	http.ListenAndServe(":8080", mux)
}
