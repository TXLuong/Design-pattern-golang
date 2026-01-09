package main

func main() {
	service := NotificationService{}
	service.notify("email", "Hello")
	service.notify("sms", "Hello")
}