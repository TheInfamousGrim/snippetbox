package main

import (
	"flag"
	"net/http"

	"github.com/charmbracelet/log"
)

type application struct{}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := &application{}

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

	format := "%s addr=%s"
	log.Infof(format, "starting server", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err.Error())
}
