package controllers

import (
	"os"
	"fmt"
	"net/smtp"

	//
	"mailtest/api/models"
)

type Server struct {
	Auth smtp.Auth
	Address string
	From string
}

func (s *Server) Initialize(senderEmail, senderPassword, smtpHost, smtpPort string){
	s.Auth = smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	s.Address = fmt.Sprintf("%s:%s", smtpHost, smtpPort)
	s.From = senderEmail
	
	//
	//test
	//

	data := models.ConfirmEmailTemplate{
		Title: "Confirm Email",
		Name: "Test",
		URL: "localhost/confirm/1234",
	}
	
	r := models.NewEmailRequest(
		[]string{os.Getenv("RECIPIENT_EMAIL")}, 
		"Test title", 
		"Hello, World!",
	)
	
	var err error
	err = r.ParseTemplate("templates/confirmEmail.gohtml", data)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = s.SendEmail(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	//
	//
}