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

	format := "%s addr=%s"
	log.Infof(format, "starting server", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	log.Fatal(err.Error())
}
