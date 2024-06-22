package main

import (
	"net/http"
	"runtime/debug"

	"github.com/charmbracelet/log"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	format := "%s method=%s uri=%s trace=%s"
	log.Errorf(format, err.Error(), method, uri, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
