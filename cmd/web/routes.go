package main

import (
    "github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/", app.home)
    r.HandleFunc("/pin/{id}", app.showPin)
    r.HandleFunc("/pin/create", app.createPin)

    return r
}