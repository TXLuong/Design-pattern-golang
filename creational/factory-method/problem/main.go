package main

func main() {
	service := NotificationService{}
	service.SendNotification("email", "Hello")
	service.SendNotification("sms","Hello")
}