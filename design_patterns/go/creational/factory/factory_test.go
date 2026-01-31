package factory

import (
	"testing"
)

func TestNotifierFactory_ReturnsEmailNotifier(t *testing.T) {
	notifier, err := NotifierFactory("email", "gmail")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	emailNotifier, ok := notifier.(*EmailNotifier)
	if !ok {
		t.Errorf("expected *EmailNotifier, got %T", notifier)
	}

	if emailNotifier.Vendor != "gmail" {
		t.Errorf("expected vendor 'gmail', got '%s'", emailNotifier.Vendor)
	}
}

func TestNotifierFactory_ReturnsSmsNotifier(t *testing.T) {
	notifier, err := NotifierFactory("sms", "twilio")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	smsNotifier, ok := notifier.(*SmsNotifier)
	if !ok {
		t.Errorf("expected *SmsNotifier, got %T", notifier)
	}

	if smsNotifier.Vendor != "twilio" {
		t.Errorf("expected vendor 'twilio', got '%s'", smsNotifier.Vendor)
	}
}

func TestNotifierFactory_ReturnsPushNotifier(t *testing.T) {
	notifier, err := NotifierFactory("push", "firebase")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	pushNotifier, ok := notifier.(*PushNotifier)
	if !ok {
		t.Errorf("expected *PushNotifier, got %T", notifier)
	}

	if pushNotifier.Vendor != "firebase" {
		t.Errorf("expected vendor 'firebase', got '%s'", pushNotifier.Vendor)
	}
}

func TestNotifierFactory_ReturnsErrorForUnknownType(t *testing.T) {
	notifier, err := NotifierFactory("unknown", "vendor")

	if err == nil {
		t.Errorf("expected error for unknown type, got nil")
	}

	if notifier != nil {
		t.Errorf("expected nil notifier for unknown type, got %T", notifier)
	}
}

func TestNotifier_PolymorphicBehavior(t *testing.T) {
	// Create different notifiers using factory
	notifiers := []struct {
		notifierType string
		vendor       string
	}{
		{"email", "gmail"},
		{"sms", "twilio"},
		{"push", "firebase"},
	}

	for _, tc := range notifiers {
		notifier, err := NotifierFactory(tc.notifierType, tc.vendor)
		if err != nil {
			t.Fatalf("unexpected error for %s: %v", tc.notifierType, err)
		}

		// All notifiers should implement Send - polymorphic call
		err = notifier.Send("test message")
		if err != nil {
			t.Errorf("Send failed for %s: %v", tc.notifierType, err)
		}
	}
}

func TestNotifier_InterfaceSatisfaction(t *testing.T) {
	// Compile-time check that all types satisfy Notifier interface
	var _ Notifier = &EmailNotifier{}
	var _ Notifier = &SmsNotifier{}
	var _ Notifier = &PushNotifier{}
}

// ==================== NEGATIVE TEST CASES ====================

func TestNotifierFactory_EmptyType(t *testing.T) {
	notifier, err := NotifierFactory("", "vendor")

	if err == nil {
		t.Errorf("expected error for empty type, got nil")
	}

	if notifier != nil {
		t.Errorf("expected nil notifier for empty type, got %T", notifier)
	}
}

func TestNotifierFactory_CaseSensitivity(t *testing.T) {
	// Factory should be case-sensitive - "EMAIL" != "email"
	testCases := []string{"EMAIL", "Email", "SMS", "Sms", "PUSH", "Push"}

	for _, tc := range testCases {
		notifier, err := NotifierFactory(tc, "vendor")
		if err == nil {
			t.Errorf("expected error for '%s' (case mismatch), got nil", tc)
		}
		if notifier != nil {
			t.Errorf("expected nil notifier for '%s', got %T", tc, notifier)
		}
	}
}

func TestNotifierFactory_SpecialCharacters(t *testing.T) {
	invalidTypes := []string{"email!", "@sms", "push#", "em ail", "sms\n", ""}

	for _, tc := range invalidTypes {
		notifier, err := NotifierFactory(tc, "vendor")
		if err == nil {
			t.Errorf("expected error for invalid type '%s', got nil", tc)
		}
		if notifier != nil {
			t.Errorf("expected nil notifier for '%s', got %T", tc, notifier)
		}
	}
}

func TestNotifierFactory_NilValueHandling(t *testing.T) {
	// Empty vendor should still work - factory doesn't validate vendor
	notifier, err := NotifierFactory("email", "")

	if err != nil {
		t.Errorf("unexpected error for empty vendor: %v", err)
	}

	if notifier == nil {
		t.Errorf("expected notifier for empty vendor, got nil")
	}

	emailNotifier, ok := notifier.(*EmailNotifier)
	if !ok {
		t.Errorf("expected *EmailNotifier, got %T", notifier)
	}

	if emailNotifier.Vendor != "" {
		t.Errorf("expected empty vendor, got '%s'", emailNotifier.Vendor)
	}
}

func TestNotifierFactory_WhitespaceType(t *testing.T) {
	whitespaceTypes := []string{" ", "  ", "\t", "\n", " email", "email "}

	for _, tc := range whitespaceTypes {
		notifier, err := NotifierFactory(tc, "vendor")
		if err == nil {
			t.Errorf("expected error for whitespace type '%q', got nil", tc)
		}
		if notifier != nil {
			t.Errorf("expected nil notifier for '%q', got %T", tc, notifier)
		}
	}
}
