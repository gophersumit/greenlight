package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "status:available")
	fmt.Fprintln(w, "version:", version)
	fmt.Fprintln(w, "env:", app.config.env)
}
