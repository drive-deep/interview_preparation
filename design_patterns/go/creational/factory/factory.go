package factory

import "fmt"

// Notifier interface - all notifiers must implement this
type Notifier interface {
	Send(message string) error
}

// EmailNotifier sends notifications via email
type EmailNotifier struct {
	Vendor string
}

// NewEmailNotifier creates a new EmailNotifier
func NewEmailNotifier(vendor string) *EmailNotifier {
	return &EmailNotifier{
		Vendor: vendor,
	}
}

// Send sends an email notification
func (e *EmailNotifier) Send(msg string) error {
	fmt.Printf("Msg '%s' sent from email vendor: %s\n", msg, e.Vendor)
	return nil
}

// SmsNotifier sends notifications via SMS
type SmsNotifier struct {
	Vendor string
}

// NewSmsNotifier creates a new SmsNotifier
func NewSmsNotifier(vendor string) *SmsNotifier {
	return &SmsNotifier{
		Vendor: vendor,
	}
}

// Send sends an SMS notification
func (e *SmsNotifier) Send(msg string) error {
	fmt.Printf("Msg '%s' sent from SMS vendor: %s\n", msg, e.Vendor)
	return nil
}

// PushNotifier sends push notifications
type PushNotifier struct {
	Vendor string
}

// NewPushNotifier creates a new PushNotifier
func NewPushNotifier(vendor string) *PushNotifier {
	return &PushNotifier{
		Vendor: vendor,
	}
}

// Send sends a push notification
func (e *PushNotifier) Send(msg string) error {
	fmt.Printf("Msg '%s' sent from Push vendor: %s\n", msg, e.Vendor)
	return nil
}

// NotifierFactory creates the appropriate Notifier based on type
func NotifierFactory(notifierType, vendor string) (Notifier, error) {
	switch notifierType {
	case "email":
		return NewEmailNotifier(vendor), nil
	case "sms":
		return NewSmsNotifier(vendor), nil
	case "push":
		return NewPushNotifier(vendor), nil
	default:
		return nil, fmt.Errorf("unknown notifier type: %s", notifierType)
	}
}