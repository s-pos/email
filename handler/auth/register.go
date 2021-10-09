package auth

import (
	"encoding/json"
	"spos/email/handler"

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
}

func NewRegisterHandler(log *logrus.Logger) *registerHandler {
	return &registerHandler{
		log: log,
	}
}

func (r *registerHandler) Handler(msg amqp.Delivery) {
	var req registerData

	err := json.Unmarshal(msg.Body, &req)
	if err != nil {
		r.log.Errorf("error when unmarshal body from queue %v", err)
		return
	}
}
