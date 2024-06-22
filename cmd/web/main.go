package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//* File Server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//* Static Assets
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	//* Home
	mux.HandleFunc("GET /{$}", home)

	//* Snippet View
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)

	//* Snippet Create
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
