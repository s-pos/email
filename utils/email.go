package utils

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"reflect"
	"strconv"

	"gopkg.in/gomail.v2"
)

type DataEmail struct {
	To            string
	Subject       string
	TemplateEmail string
	Data          map[string]interface{}
}

func (d *DataEmail) SendEmail() error {
	var (
		emailSender = os.Getenv("CONFIG_AUTH_EMAIL")
		password    = os.Getenv("CONFIG_AUTH_PASSWORD")
		host        = os.Getenv("CONFIG_SMTP_HOST")
		port, _     = strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
		sender      = fmt.Sprintf("%s<%s>", os.Getenv("CONFIG_SENDER_NAME"), emailSender)
		err         error
	)
	if reflect.ValueOf(d.Data).IsZero() {
		err = errors.New("data not found")
		log.Print(err)
		return err
	}

	buffer := new(bytes.Buffer)
	t, err := template.ParseFiles(d.TemplateEmail)
	if err != nil {
		log.Print(err)
		err = fmt.Errorf("error parse file %v", err)
		return err
	}

	err = t.Execute(buffer, d.Data)
	if err != nil {
		log.Print(err)
		err = fmt.Errorf("error execute template %v", err)
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", sender)
	mailer.SetHeader("To", d.To)
	mailer.SetHeader("Subject", d.Subject)
	mailer.SetBody("text/html", buffer.String())

	dialer := gomail.NewDialer(
		host,
		port,
		emailSender,
		password,
	)

	err = dialer.DialAndSend(mailer)
	log.Print(err)
	return err
}
