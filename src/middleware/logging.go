package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Start %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index handler")
	fmt.Fprintf(w, "welcome")
}
func icoHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	http.HandleFunc("/favicon.ico", icoHandler)
	indexHandler := http.HandlerFunc(index)
	http.Handle("/", loggingHandler(indexHandler))
	server := &http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
