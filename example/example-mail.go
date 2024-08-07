package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	ExamplePlainAuth()
}

func Example() {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("mxa-0001b201.gslb.pphosted.com:25")
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("no-reply@yourdomain.com"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("k.zhang@dayforce.com"); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}

// variables to make ExamplePlainAuth compile, without adding
// unnecessary noise there.
var (
	from       = "no-reply@yourdomain.com"
	msg        = []byte("dummy message")
	recipients = []string{"k.zhang@ceridian.com"}
)

func ExamplePlainAuth() {
	// hostname is used by PlainAuth to validate the TLS certificate.
	hostname := "mxa-0001b201.gslb.pphosted.com"
	//auth := smtp.PlainAuth("", "user@example.com", "password", hostname)

	err := smtp.SendMail(hostname+":25", nil, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleSendMail() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"recipient@example.net"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
