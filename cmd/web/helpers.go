package main

import (
	"fmt"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	app.errorLog.Printf(fmt.Sprintf("%s\n", err.Error()))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
