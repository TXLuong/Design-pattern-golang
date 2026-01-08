package  main 

import "fmt"

type NotificationService struct {
	notifierName string
}

func (s NotificationService) SendNotification(message string) {
	if s.notifierName == "email" {
		fmt.Printf("Send message: %s (Sender: Email) \n", message)
	} else if (s.notifierName == "sms") {
		fmt.Printf("Send message: %s (Sender: SMS) \n", message)
	}
}

func main() {
	s := NotificationService{notifierName : "email"}
	s.SendNotification("Hello, Send with problem of if else")
}