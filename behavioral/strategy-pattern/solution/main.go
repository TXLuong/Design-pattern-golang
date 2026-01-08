package main

import "fmt"



func main() {
	// context 
	message := "Sebastian's learning strategy pattern"
	emailSender := EmailSender{}
	smsSender := SmsSender{}
	service := EmailService{}
	// send via email 
	fmt.Printf("Sending message via email \n")
	service.setSender(emailSender)
	service.sendMessage(message)

	// send via sms 
	fmt.Printf("Sending message via sms \n")
	service.setSender(smsSender)
	service.sendMessage(message)


}