package main

import (
	"forum/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", handlers.Index)

	svr := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on port :8080")
	svr.ListenAndServe()
}
