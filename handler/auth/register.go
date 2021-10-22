package auth

import (
	"encoding/json"
	"spos/email/handler"
	"spos/email/utils"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type registerHandler struct {
	handler.BaseHandler
	log *logrus.Logger
}

type registerData struct {
	OTP   string `json:"otp"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Link  string `json:"link"`
}

func NewRegisterHandler(log *logrus.Logger) *registerHandler {
	return &registerHandler{
		log: log,
	}
}

func (r *registerHandler) Handler(msg amqp.Delivery) {
	var (
		req registerData
	)

	err := json.Unmarshal(msg.Body, &req)
	if err != nil {
		r.log.Errorf("error when unmarshal body from queue %v", err)
		return
	}

	data := map[string]interface{}{
		"otp":  req.OTP,
		"link": req.Link,
		"name": req.Name,
	}

	dataEmail := &utils.DataEmail{
		To:            req.Email,
		Subject:       "Selamat datang di SPOS",
		TemplateEmail: "templates/register.html",
		Data:          data,
	}

	err = dataEmail.SendEmail()
	if err != nil {
		r.log.Error(err)
	} else {
		r.log.Infof("success send email registration %s", req.Email)
	}
}
