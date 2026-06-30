package main

import (
	"fmt"
)

type NotificationSender interface{
	send (message string)
}

type EmailSender struct{}

type (e *EmailSender) send(message string){
	fmt.Println("Welcome" + message)
}

type SmsSender struct{}

type (e *SmsSender) send(message string){
	fmt.Println("Welcome" + message)
}

type NotificationService struct {
	sender NotificationSender
}

func main(){
	fmt.Println("Welcome hello world")

	email := &EmailSender{}

	service :=&NotificationService{
		sender: email,
	}

	service.sender.send("Pratheep")

}