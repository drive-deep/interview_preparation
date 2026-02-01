package abstract_factory

import "fmt"

// ============================================================================
// PRODUCT INTERFACES - Define what each product type can do
// ============================================================================

// Button represents a clickable button component
type Button interface {
	Render() string
	OnClick(handler func())
}

// Checkbox represents a toggleable checkbox component
type Checkbox interface {
	Render() string
	Toggle()
	IsChecked() bool
}

// ============================================================================
// ABSTRACT FACTORY - Creates families of related products
// ============================================================================

// GUIFactory creates UI components for a specific platform
type GUIFactory interface {
	CreateButton(label string) Button
	CreateCheckbox(label string) Checkbox
}

// ============================================================================
// WINDOWS PRODUCT FAMILY
// ============================================================================

// WindowsButton is a Windows-styled button
type WindowsButton struct {
	label   string
	onClick func()
}

func (b *WindowsButton) Render() string {
	return fmt.Sprintf("[Windows Button: %s]", b.label)
}

func (b *WindowsButton) OnClick(handler func()) {
	b.onClick = handler
}

// WindowsCheckbox is a Windows-styled checkbox
type WindowsCheckbox struct {
	label   string
	checked bool
}

func (c *WindowsCheckbox) Render() string {
	check := "☐"
	if c.checked {
		check = "☑"
	}
	return fmt.Sprintf("[Windows %s %s]", check, c.label)
}

func (c *WindowsCheckbox) Toggle() {
	c.checked = !c.checked
}

func (c *WindowsCheckbox) IsChecked() bool {
	return c.checked
}

// WindowsFactory creates Windows UI components
type WindowsFactory struct{}

func NewWindowsFactory() *WindowsFactory {
	return &WindowsFactory{}
}

func (f *WindowsFactory) CreateButton(label string) Button {
	return &WindowsButton{label: label}
}

func (f *WindowsFactory) CreateCheckbox(label string) Checkbox {
	return &WindowsCheckbox{label: label}
}

// ============================================================================
// MACOS PRODUCT FAMILY
// ============================================================================

// MacOSButton is a macOS-styled button
type MacOSButton struct {
	label   string
	onClick func()
}

func (b *MacOSButton) Render() string {
	return fmt.Sprintf("( macOS Button: %s )", b.label)
}

func (b *MacOSButton) OnClick(handler func()) {
	b.onClick = handler
}

// MacOSCheckbox is a macOS-styled checkbox
type MacOSCheckbox struct {
	label   string
	checked bool
}

func (c *MacOSCheckbox) Render() string {
	check := "○"
	if c.checked {
		check = "●"
	}
	return fmt.Sprintf("( macOS %s %s )", check, c.label)
}

func (c *MacOSCheckbox) Toggle() {
	c.checked = !c.checked
}

func (c *MacOSCheckbox) IsChecked() bool {
	return c.checked
}

// MacOSFactory creates macOS UI components
type MacOSFactory struct{}

func NewMacOSFactory() *MacOSFactory {
	return &MacOSFactory{}
}

func (f *MacOSFactory) CreateButton(label string) Button {
	return &MacOSButton{label: label}
}

func (f *MacOSFactory) CreateCheckbox(label string) Checkbox {
	return &MacOSCheckbox{label: label}
}

// ============================================================================
// CLIENT CODE - Works with factories and products via interfaces only
// ============================================================================

// Application uses abstract factory to create UI - doesn't know concrete types
type Application struct {
	factory  GUIFactory
	button   Button
	checkbox Checkbox
}

// NewApplication creates an app with the given UI factory
func NewApplication(factory GUIFactory) *Application {
	return &Application{factory: factory}
}

// CreateUI builds the user interface using the factory
func (a *Application) CreateUI() {
	a.button = a.factory.CreateButton("Submit")
	a.checkbox = a.factory.CreateCheckbox("Remember me")
}

// RenderUI returns the rendered UI components
func (a *Application) RenderUI() string {
	return fmt.Sprintf("%s\n%s", a.button.Render(), a.checkbox.Render())
}

// ============================================================================
// FACTORY SELECTOR - Returns appropriate factory based on config/environment
// ============================================================================

// GetGUIFactory returns the appropriate factory for the given OS
func GetGUIFactory(osType string) GUIFactory {
	switch osType {
	case "windows":
		return NewWindowsFactory()
	case "macos":
		return NewMacOSFactory()
	default:
		// Default to Windows
		return NewWindowsFactory()
	}
}
