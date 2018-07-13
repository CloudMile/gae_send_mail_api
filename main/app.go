package app

import (
	"net/http"

	"github.com/CloudMile/gae_send_mail_api/controller"
)

func init() {
	router := controller.Router()
	http.Handle("/", router)
}
