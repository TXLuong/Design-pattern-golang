package main


type EmailService struct {
	sender Sender
}

func (service *EmailService) setSender(sender Sender) {
	service.sender = sender
}

func (service *EmailService) sendMessage(message string) {
	service.sender.Send(message)
}