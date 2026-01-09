package main

import "fmt"

type EmailNotification struct{}

func (e EmailNotification) Send(content string) {
	fmt.Printf("Send message %s via Email \n", content)
}