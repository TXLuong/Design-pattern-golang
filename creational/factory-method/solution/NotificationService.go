package main 

type NotificationService struct {

}

func (n NotificationService) notify(notificationType string, content string) {
	// Use factory method 
	notifier := NewNotification(notificationType)
	notifier.Send(content)
}