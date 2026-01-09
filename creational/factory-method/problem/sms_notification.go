package main

import "fmt"

type SmsNotification struct{}

func (s SmsNotification) Send(message string) {
fmt.Printf("Send message %s via SMS \n", message)
}