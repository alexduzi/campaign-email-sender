package mailer

import (
	"campaignemailsender/internal/domain/campaign"
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(campaign *campaign.Campaign) (err error) {
	log.Println("Sending Email...")

	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	var emails []string
	for _, contact := range campaign.Contacts {
		emails = append(emails, contact.Email)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", emails...)
	m.SetHeader("Subject", fmt.Sprintf("Campaign email sender! -> %v", campaign.Name))
	m.SetBody("text/html", campaign.Content)

	err = d.DialAndSend(m)

	return
}
