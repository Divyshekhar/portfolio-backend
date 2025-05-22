package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type ContactForm struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func MailController(c *gin.Context) {

	var form ContactForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	SendEmail(form)
}

func SendEmail(form ContactForm) error {

	emailTo := os.Getenv("EMAIL_TO")
	emailFrom := os.Getenv("EMAIL_FROM")
	pass := os.Getenv("EMAIL_PASS")
	m := gomail.NewMessage()

	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", "New Contact Form Submission: "+form.Subject)

	body := "You have a new message from your website:\n\n" +
		"Name: " + form.Name + "\n" +
		"Email: " + form.Email + "\n" +
		"Subject: " + form.Subject + "\n" +
		"Message:\n" + form.Message

	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, emailFrom, pass)
	return d.DialAndSend(m)

}
