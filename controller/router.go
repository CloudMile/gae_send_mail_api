package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

// Router is router list
func Router() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/send", AddContext(Send)).Methods("POST")

	return
}

// AddContext is modify http.Request, cintext for need
func AddContext(handleFunc func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := GetVars(r)
		ctx := appengine.NewContext(r)
		r = mux.SetURLVars(r.WithContext(ctx), vars)
		handleFunc(w, r)
	}
}
