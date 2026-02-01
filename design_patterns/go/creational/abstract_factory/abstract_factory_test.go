package abstract_factory

import (
	"strings"
	"testing"
)

func TestWindowsFactory_CreateButton(t *testing.T) {
	factory := NewWindowsFactory()
	button := factory.CreateButton("Click Me")

	rendered := button.Render()
	if !strings.Contains(rendered, "Windows") {
		t.Errorf("Expected Windows button, got: %s", rendered)
	}
	if !strings.Contains(rendered, "Click Me") {
		t.Errorf("Expected label 'Click Me', got: %s", rendered)
	}
}

func TestWindowsFactory_CreateCheckbox(t *testing.T) {
	factory := NewWindowsFactory()
	checkbox := factory.CreateCheckbox("Accept Terms")

	// Initially unchecked
	if checkbox.IsChecked() {
		t.Error("Checkbox should be unchecked initially")
	}

	// Toggle and verify
	checkbox.Toggle()
	if !checkbox.IsChecked() {
		t.Error("Checkbox should be checked after toggle")
	}

	rendered := checkbox.Render()
	if !strings.Contains(rendered, "Windows") {
		t.Errorf("Expected Windows checkbox, got: %s", rendered)
	}
}

func TestMacOSFactory_CreateButton(t *testing.T) {
	factory := NewMacOSFactory()
	button := factory.CreateButton("Submit")

	rendered := button.Render()
	if !strings.Contains(rendered, "macOS") {
		t.Errorf("Expected macOS button, got: %s", rendered)
	}
}

func TestMacOSFactory_CreateCheckbox(t *testing.T) {
	factory := NewMacOSFactory()
	checkbox := factory.CreateCheckbox("Remember")

	rendered := checkbox.Render()
	if !strings.Contains(rendered, "macOS") {
		t.Errorf("Expected macOS checkbox, got: %s", rendered)
	}
}

func TestApplication_WithWindowsFactory(t *testing.T) {
	app := NewApplication(NewWindowsFactory())
	app.CreateUI()

	ui := app.RenderUI()
	if !strings.Contains(ui, "Windows Button") {
		t.Errorf("Expected Windows UI, got: %s", ui)
	}
}

func TestApplication_WithMacOSFactory(t *testing.T) {
	app := NewApplication(NewMacOSFactory())
	app.CreateUI()

	ui := app.RenderUI()
	if !strings.Contains(ui, "macOS Button") {
		t.Errorf("Expected macOS UI, got: %s", ui)
	}
}

func TestGetGUIFactory(t *testing.T) {
	tests := []struct {
		os       string
		expected string
	}{
		{"windows", "Windows"},
		{"macos", "macOS"},
		{"linux", "Windows"}, // defaults to Windows
	}

	for _, tt := range tests {
		factory := GetGUIFactory(tt.os)
		button := factory.CreateButton("Test")
		if !strings.Contains(button.Render(), tt.expected) {
			t.Errorf("GetGUIFactory(%s): expected %s, got %s", tt.os, tt.expected, button.Render())
		}
	}
}

// TestFactoryProducesConsistentFamily verifies that all products from
// a factory belong to the same family (key benefit of Abstract Factory)
func TestFactoryProducesConsistentFamily(t *testing.T) {
	winFactory := NewWindowsFactory()
	winButton := winFactory.CreateButton("Btn")
	winCheckbox := winFactory.CreateCheckbox("Chk")

	// Both should be Windows-styled
	if !strings.Contains(winButton.Render(), "Windows") ||
		!strings.Contains(winCheckbox.Render(), "Windows") {
		t.Error("Windows factory should produce consistent Windows-styled components")
	}

	macFactory := NewMacOSFactory()
	macButton := macFactory.CreateButton("Btn")
	macCheckbox := macFactory.CreateCheckbox("Chk")

	// Both should be macOS-styled
	if !strings.Contains(macButton.Render(), "macOS") ||
		!strings.Contains(macCheckbox.Render(), "macOS") {
		t.Error("macOS factory should produce consistent macOS-styled components")
	}
}
