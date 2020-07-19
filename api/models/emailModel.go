package models

import (
	"bytes"
	"html/template"
)

type EmailRequest struct {
	From    string
	To      []string
	Subject string
	Body    string
}

type ConfirmEmailTemplate struct {
	Title string
	Name string
	URL string
}

func NewEmailRequest(to []string, subject, body string) *EmailRequest {
	return &EmailRequest{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func (r *EmailRequest) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.Body = buf.String()
	return nil
}