package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	//* File Server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	//* Static Assets
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	//* Home
	mux.HandleFunc("GET /{$}", app.home)

	//* Snippet View
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)

	//* Snippet Create
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
