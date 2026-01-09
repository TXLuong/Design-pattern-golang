package main

type NotificationService struct {
	name string
}

// Notification Service will depend on concrete Notifier 
// When we have new notification method -> we need to change in this function -> violate OPEN/CLOSE principal 
func (n *NotificationService) SendNotification(notificationType string, content string) {
	switch notificationType {
	case "email":
		email := EmailNotification{}
		email.Send(content)
	case "sms":
		sms := SmsNotification{}
		sms.Send(content)
	}
}