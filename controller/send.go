package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/CloudMile/gae_send_mail_api/model"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
)

// HeaderContentType is what Content-Type this API use
var HeaderContentType = map[string]map[string]bool{
	"multipart/form-data":               {"pass": true, "isDataUpload": true},
	"application/x-www-form-urlencoded": {"pass": true, "isDataUpload": false},
}

// Send is the an endpoint "POST /send"
func Send(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("CUSTOM_TOKEN") != "" && r.Header.Get("custom-token") != os.Getenv("CUSTOM_TOKEN") {
		ErrorResponse(w, r, http.StatusNonAuthoritativeInfo, nil, "auth wrong")
		return
	}
	ctx := r.Context()
	log.Infof(ctx, "POST /send")

	ct := r.Header.Get("Content-Type")
	log.Infof(ctx, "Content-Type is: %s", ct)

	contentType, pass, isDataUpload := checkContentType(ctx, ct)
	if !pass {
		ErrorResponse(w, r, http.StatusNonAuthoritativeInfo, nil, "Content-Type error, shoud be multipart/form-data or application/x-www-form-urlencoded")
		return
	}
	form := MakeMailParams(r, contentType)
	log.Infof(ctx, "url values is %+v", form)

	w.Header().Set("Content-Type", "application/json")
	log.Infof(ctx, "send to %s", form.To)
	log.Infof(ctx, "isDataUpload is %v", isDataUpload)

	var attachments []mail.Attachment
	if isDataUpload {
		createdAttachments, err := CreateAttachments(r)
		attachments = createdAttachments
		if err != nil {
			ErrorResponse(w, r, http.StatusUnprocessableEntity, err, "upload file failed")
		}
	}

	gaeMail := MakeGaeMail(ctx, &form, attachments)
	sendErr := gaeMail.Send()
	if sendErr != nil {
		ErrorResponse(w, r, http.StatusUnprocessableEntity, sendErr, "send mail failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", `{"result": "sent success"}`)
}

// CreateAttachments is create attachment from r.FormFile
func CreateAttachments(r *http.Request) (attachments []mail.Attachment, err error) {
	file, header, err := r.FormFile("data")
	if err != nil && err.Error() != `http: no such file` {
		return
	}
	attachments, err = MakeAttachments(r, file, header)
	if err != nil {
		return
	}
	return
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

// MakeMailParams is make mail params
func MakeMailParams(r *http.Request, contentType string) (form model.Form) {
	switch contentType {
	case "multipart/form-data":
		form = model.Form{
			To:      r.FormValue("to"),
			CC:      r.FormValue("cc"),
			BCC:     r.FormValue("bcc"),
			Subject: r.FormValue("subject"),
			Body:    r.FormValue("body"),
		}
	case "application/x-www-form-urlencoded":
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &form)
	}
	return
}

// MakeGaeMail is make model GaeMail
func MakeGaeMail(ctx context.Context, form *model.Form, attachments []mail.Attachment) (gaeMail model.GaeMail) {
	gaeMail = model.GaeMail{
		Ctx:     ctx,
		To:      form.To,
		CC:      form.CC,
		BCC:     form.BCC,
		Subject: form.Subject,
		Body:    form.Body,
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

func checkContentType(ctx context.Context, contentType string) (reContentType string, pass, isDataUpload bool) {
	reContentType = strings.Split(contentType, ";")[0]
	log.Infof(ctx, "split contentType is: %s", reContentType)

	headerContentType := HeaderContentType[reContentType]
	pass = headerContentType["pass"]
	isDataUpload = headerContentType["isDataUpload"]
	return
}
