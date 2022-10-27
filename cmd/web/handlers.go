package main

import (
    "errors"
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"
    
    "github.com/Krokin/PinBox/pkg/models"

    "github.com/gorilla/mux"
)

func (app *application) response(w http.ResponseWriter, p[] *models.Pin) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(p)
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }
    p, err := app.pins.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }
    app.response(w, p)

}

func (app *application) showPin(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }
    var pins []*models.Pin
    p, err := app.pins.Get(id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecord) {
            app.notFound(w)
        } else {
            app.serverError(w, err)
        }
        return
    }
    pins = append(pins, p)
    app.response(w, pins)
}

func (app *application) createPin(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }
    p := models.Pin{}
    json.NewDecoder(r.Body).Decode(&p)
    id, err := app.pins.Insert(p.Title, p.Content, "7")
    if err != nil {
        app.serverError(w, err)
        return
    }
    http.Redirect(w, r, fmt.Sprintf("/pin/%d", id), http.StatusSeeOther)
}