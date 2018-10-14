package controller

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/CloudMile/gae_send_mail_api/model"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
)

// Send is the an endpoint "POST /send"
func Send(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("CUSTOM_TOKEN") != "" && r.Header.Get("custom-token") != os.Getenv("CUSTOM_TOKEN") {
		ErrorResponse(w, r, http.StatusNonAuthoritativeInfo, nil, "auth wrong")
		return
	}
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	log.Infof(ctx, "POST /send")
	log.Infof(ctx, "send to %s", r.FormValue("to"))

	file, header, err := r.FormFile("data")
	if err != nil && err.Error() != `http: no such file` {
		ErrorResponse(w, r, http.StatusUnprocessableEntity, err, "upload file failed")
		return
	}

	attachments, chErr := MakeAttachments(r, file, header)
	if chErr != nil {
		ErrorResponse(w, r, http.StatusUnprocessableEntity, chErr, "upload file failed")
		return
	}

	gaeMail := MakeGaeMail(r, attachments)
	sendErr := gaeMail.Send()
	if sendErr != nil {
		ErrorResponse(w, r, http.StatusUnprocessableEntity, sendErr, "send mail failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", `{"result": "sent success"}`)
}

// MakeAttachments is using model UploadToAttachment to create attachment
func MakeAttachments(r *http.Request, file multipart.File, header *multipart.FileHeader) (attachments []mail.Attachment, err error) {
	attachments = make([]mail.Attachment, 0)

	if file != nil {
		upload := model.UploadToAttachment{
			UploadFile:   file,
			UploadHeader: header,
		}
		err = upload.Change()

		if err != nil {
			return
		}
		attachments = append(attachments, upload.Attachment)
	}
	return
}

// MakeGaeMail is make model GaeMail
func MakeGaeMail(r *http.Request, attachments []mail.Attachment) (gaeMail model.GaeMail) {
	gaeMail = model.GaeMail{
		Ctx:     r.Context(),
		To:      r.FormValue("to"),
		CC:      r.FormValue("cc"),
		BCC:     r.FormValue("bcc"),
		Subject: r.FormValue("subject"),
		Body:    r.FormValue("body"),
	}

	if len(attachments) > 0 {
		gaeMail.Attachments = attachments
	}
	return
}

// ErrorResponse to return failed action
func ErrorResponse(w http.ResponseWriter, r *http.Request, httpStatus int, err error, errorMessage string) {
	log.Errorf(r.Context(), "Error is %s", err)
	w.WriteHeader(httpStatus)
	fmt.Fprintf(w, "%s", `{"error": "`+errorMessage+`"}`)
	return
}
