package main

import (
	"flag"
	"net/http"

	"github.com/charmbracelet/log"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

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

	format := "%s addr=%s"
	log.Infof(format, "starting server", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err.Error())
}
