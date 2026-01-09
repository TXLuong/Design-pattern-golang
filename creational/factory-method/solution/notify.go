package main 

type Notifier interface {
	Send(content string)
}

func NewNotification(notificationType string) Notifier {
	switch notificationType {
	case "email" : 
		return EmailNotification{}
	
	case "sms" : 
		return SmsNotification{}
	default : 
		return nil
	}
}


