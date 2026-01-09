package main

type NotificationService struct {}

// Notification Service will depend on concrete Notifier 
// When we have new notification method -> we need to change in this function -> violate OPEN/CLOSE principal 
// 
func (n NotificationService) SendNotification(notificationType string, content string) {
	if notificationType == "email" {
		email := EmailNotification{}
		email.Send(content)
	} else if notificationType == "sms" {
		sms := SmsNotification{}
		sms.Send(content)
	}
}