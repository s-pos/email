package utils

import (
	"os"
	"testing"
)

func init() {
	s := map[string]string{
		"CONFIG_AUTH_EMAIL":    "panduudp@gmail.com",
		"CONFIG_AUTH_PASSWORD": "Oranggila26797",
		"CONFIG_SMTP_HOST":     "smtp.gmail.com",
		"CONFIG_SMTP_PORT":     "465",
		"CONFIG_SENDER_NAME":   "no-reply",
	}

	for key, val := range s {
		os.Setenv(key, val)
	}
}

func TestSendEmail(t *testing.T) {
	s := []struct {
		name string
		data *DataEmail
	}{
		{
			name: "test case 1",
			data: &DataEmail{
				To:      "pandudpn@yopmail.com",
				Subject: "Selamat datang di SPOS",
				Data: map[string]interface{}{
					"otp":  "123456",
					"link": "https://www.google.com",
					"name": "pandu dwi putra",
				},
				TemplateEmail: "../templates/register.html",
			},
		},
	}

	for _, tt := range s {
		t.Run(tt.name, func(t *testing.T) {
			tt.data.SendEmail()
		})
	}
}
