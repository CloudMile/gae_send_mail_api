package main

import (
	"net/http"

	"github.com/CloudMile/gae_send_mail_api/controller"
	"google.golang.org/appengine/v2"
)

func main() {
	router := controller.Router()
	http.Handle("/", router)
	appengine.Main()
}
