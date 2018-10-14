package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GetVars for creating route variables from mux
func GetVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}
