package main

import "fmt"

type SmsNotification struct{}

func (s SmsNotification) Send(content string) {
	fmt.Printf("Send message %s via Sms \n", content)
}