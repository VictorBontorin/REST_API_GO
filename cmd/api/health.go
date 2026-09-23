package main

import "net/http"

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
