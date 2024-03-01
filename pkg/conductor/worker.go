package cond

import (
	"log"

	"gopkg.in/gomail.v2"
)

func (c *Conductor) worker(emailChan chan gomail.Message) {

	var (
		mail gomail.Message
		err  error
	)

	defer close(emailChan)

	for {
		select {
		case mail = <-emailChan:
			if err = c.send(&mail); err != nil {
				log.Println(err)
			}
		}
	}
}

func (c *Conductor) send(mail *gomail.Message) error {

	var (
		conn *gomail.Dialer
	)

	// Make connection to the SMTP server
	conn = gomail.NewDialer("smtp.gmail.com", c.config.SenderPort, c.config.SenderEmail, c.config.SenderPass)

	// Send the mail
	return conn.DialAndSend(mail)
}
