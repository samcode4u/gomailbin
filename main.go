package main

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var (
	SmtpUser  = os.Args[1]
	SmtpPass  = os.Args[2]
	Host      = os.Args[3]
	Port      = os.Args[4]
	CharSet   = "UTF-8"
	FromEmail = os.Args[5]
	ToEmail   = os.Args[6]
)

func SendTestMail(sender string, senderName string, recipient string, recipientName string) error {

	// Create a new message.
	m := gomail.NewMessage()

	var emailHTMLBody string

	emailHTMLBody = "Hello World Test : From " + SmtpUser

	// Set the main email part to use HTML.
	m.SetBody("text/html", emailHTMLBody)

	// Construct the message headers, including a Configuration Set and a Tag.
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(sender, senderName)},
		"To":      {recipient},
		"Subject": {"Hello World Test"},
		// Comment or remove the next line if you are not using a configuration set
		// "X-SES-CONFIGURATION-SET": {ConfigSet},
		// Comment or remove the next line if you are not using custom tags
		// "X-SES-MESSAGE-TAGS": {Tags},
	})

	m.SetAddressHeader("Cc", os.Args[1], "Team")

	port, _ := strconv.Atoi(Port)
	// Send the email.
	d := gomail.NewPlainDialer(Host, port, SmtpUser, SmtpPass)

	// Display an error message if something goes wrong; otherwise,
	// display a message confirming that the message was sent.
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func main() {
	err := SendTestMail(FromEmail, "SMTPTestFrom", ToEmail, "SMTPTestTo")
	fmt.Println(err)
}
