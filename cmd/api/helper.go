package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}
	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	w.Write(res)
	return nil
}
