package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	app.writeJSON(w, http.StatusOK, data, nil)
}
