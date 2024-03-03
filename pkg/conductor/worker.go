package cond

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/gomail.v2"
)

type messageInfo struct {
	email string
	code  string
}

func (c *Conductor) worker(emailChan chan messageInfo) {

	var (
		info    messageInfo
		message *gomail.Message
		err     error
	)

	defer close(emailChan)

	for {
		select {
		case info = <-emailChan:
			message = c.makeMessage(info.email, info.code)
			if err = c.send(message); err != nil {
				log.Println(err)
			}
		}
	}
}

func (c *Conductor) send(message *gomail.Message) error {

	var (
		conn *gomail.Dialer
	)

	// Make connection to the SMTP server
	conn = gomail.NewDialer("smtp.gmail.com", c.config.SenderPort, c.config.SenderEmail, c.config.SenderPass)

	// Send the mail
	return conn.DialAndSend(message)
}

func (c *Conductor) makeMessage(email string, verificationCode string) *gomail.Message {

	var (
		text    string = c.mailTemplate
		message *gomail.Message
	)

	text = strings.Replace(text, "{{.VerificationCode}}", verificationCode, 1)
	text = strings.Replace(text, "{{.VerificationDuration}}", fmt.Sprintf("%d", c.config.MailCodeDuration), 1)
	message = gomail.NewMessage()
	message.SetHeader("From", c.config.SenderEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", c.config.MailTitle)
	message.SetBody("text/html", text)
	return message
}
