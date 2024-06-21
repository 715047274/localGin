package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/smtp"
)

func main() {

	//var (
	//	mx  string
	//	err error
	//)
	//msg := []byte("To: k.zhang@ceridian.com\r\n" +
	//	"Subject: Hi there!\r\n" +
	//	"Content-Type: text/plain; charset=UTF-8\r\n" +
	//	"\r\n" +
	//	"Hi!\r\n")
	//to := make([]string, 1)
	//to[0] = "autotest@dayforce.com"
	//
	//mx, err = getMXRecord("corpadds.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//res := smtp.SendMail(mx+":25",
	//	nil, /* this is the optional Auth */
	//	"autotest@dayforce.com", to, msg)
	//if res != nil {
	//	fmt.Println(res)
	//}
	//log.Fatal(res)

	///////////////////////////////////////

	// Connect to the remote SMTP server.
	var (
		mx  string
		err error
	)
	mx, err = getMXRecord("corpadds.com")
	if err != nil {
		log.Fatal(err)
	}
	c, err := smtp.Dial(mx + ":25")
	if err != nil {
		log.Fatal(err)
	}
	// Set the sender and recipient.
	c.Mail("autotest@dayforce.com")
	c.Rcpt("k.zhang@dayforce.com")
	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("This is the email body.")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}

}

func getMXRecord(to string) (mx string, err error) {
	// var e *mail.Address
	//e, err = mail.ParseAddress(to)
	//if err != nil {
	//	return
	//}

	domain := to

	var mxs []*net.MX
	mxs, err = net.LookupMX(domain)

	if err != nil {
		return
	}

	for _, x := range mxs {
		mx = x.Host
		fmt.Println(mx)
		return
	}

	return
}
