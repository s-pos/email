package main

import (
	handlerAuth "spos/email/handler/auth"
	queueAuth "spos/email/queue/auth"

	"github.com/s-pos/go-utils/adapter"
	"github.com/s-pos/go-utils/config"
)

func main() {
	log := config.Logrus()
	channel := adapter.ConnectionAMQP()

	registerQueue := queueAuth.NewRegisterQueue(channel)
	registerQueue.Setup()
	register := handlerAuth.NewRegisterHandler(log)
	registerQueue.HandleConsume(register.Handler)

	log.Println("Service Email is Running")
	forever := make(chan bool)
	<-forever
}
