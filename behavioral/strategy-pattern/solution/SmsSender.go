package main

import "fmt"

type SmsSender struct{}

func (s SmsSender) Send(message string) {
	fmt.Printf("Message %s is sent from SMS sender \n ", message)
}

