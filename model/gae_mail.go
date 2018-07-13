package model

import (
	"context"
	"os"
	"strings"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
)

// GaeMail store the mail infos
type GaeMail struct {
	Ctx         context.Context
	To          string `json:"to"`
	CC          string `json:"cc"`
	BCC         string `json:"bcc"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	Attachments []mail.Attachment
}

// Send will send mail
func (gaeMail *GaeMail) Send() (err error) {
	ctx := gaeMail.Ctx
	msg := &mail.Message{
		Sender:  os.Getenv("MAIL_SENDER"),
		To:      strings.Split(gaeMail.To, ","),
		Cc:      strings.Split(gaeMail.CC, ","),
		Bcc:     strings.Split(gaeMail.BCC, ","),
		Subject: gaeMail.Subject,
		Body:    gaeMail.Body,
	}

	if len(gaeMail.Attachments) > 0 {
		msg.Attachments = gaeMail.Attachments
		log.Infof(ctx, "Has attachment")
	}

	if err = mail.Send(ctx, msg); err != nil {
		log.Errorf(ctx, "Couldn't send email: %v", err)
	} else {
		log.Infof(ctx, "Mail send to %s", gaeMail.To)
	}
	return
}
