package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func Router() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/send", AddContext(Send)).Methods("POST")

	return
}

func AddContext(handleFunc func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := GetVars(r)
		ctx := appengine.NewContext(r)
		r = mux.SetURLVars(r.WithContext(ctx), vars)
		handleFunc(w, r)
	}
}
