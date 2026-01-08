package main
import "fmt"

type EmailSender struct {}

func (e EmailSender) Send(message string) {
	fmt.Printf("Send %s from Email Sender \n", message)
}