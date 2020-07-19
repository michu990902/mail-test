package controllers

import (
	"net/smtp"

	"mailtest/api/models"
)

func (s *Server) SendEmail(r *models.EmailRequest) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.Body)

	return smtp.SendMail(s.Address, s.Auth, s.From, r.To, msg)
}