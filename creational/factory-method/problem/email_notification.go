package main

import "fmt"

type EmailNotification struct{}

func (e EmailNotification) Send(message string) {
	fmt.Printf("Send message: %s via Email \n", message)
}