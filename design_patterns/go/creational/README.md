# Creational Design Patterns in Go

## What Are Creational Patterns?

Creational patterns deal with **HOW objects are created**. They abstract the instantiation process, making systems independent of how objects are created, composed, and represented.

```
┌─────────────────────────────────────────────────────────────────┐
│                    CREATIONAL PATTERNS                          │
│                                                                 │
│   "Control object creation to make your system more flexible"   │
│                                                                 │
│   Instead of:  obj := &MyStruct{}                              │
│   You get:     obj := Factory.Create() or Builder.Build()      │
└─────────────────────────────────────────────────────────────────┘
```

## When to Use Creational Patterns

| Situation | Pattern to Use |
|-----------|----------------|
| Need exactly ONE instance globally | **Singleton** |
| Don't know exact type until runtime | **Factory** |
| Object has many optional parameters | **Builder** |
| Need to clone existing objects | **Prototype** |
| Create families of related objects | **Abstract Factory** |

---

## The 5 Creational Patterns

### 1. Singleton
**Problem**: Ensure only ONE instance of a class exists globally.

**Go Implementation**: Use `sync.Once` for thread-safe lazy initialization.

```go
var (
    instance *Database
    once     sync.Once
)

func GetInstance() *Database {
    once.Do(func() {
        instance = &Database{connection: "connected"}
    })
    return instance
}
```

**Real Examples**: Logger, Config Manager, Database Connection Pool, Cache

**Interview Tip**: Mention thread-safety concerns and `sync.Once`

---

### 2. Factory Method
**Problem**: Create objects without specifying the exact concrete class.

**Go Implementation**: Function that returns an interface type.

```go
type Notifier interface {
    Send(message string) error
}

func NewNotifier(notifierType string) Notifier {
    switch notifierType {
    case "email":
        return &EmailNotifier{}
    case "sms":
        return &SMSNotifier{}
    default:
        return &PushNotifier{}
    }
}
```

**Real Examples**: Database drivers, Payment gateways, Parser factories

**Interview Tip**: Client code depends on interface, not concrete types

---

### 3. Builder
**Problem**: Construct complex objects step-by-step with many optional parameters.

**Go Implementation**: Fluent interface with method chaining OR functional options.

```go
// Fluent Builder
server := NewServerBuilder().
    WithHost("api.example.com").
    WithPort(443).
    WithTLS(true).
    WithTimeout(30 * time.Second).
    Build()

// Functional Options (more Go-idiomatic)
server := NewServer(
    WithHost("api.example.com"),
    WithPort(443),
    WithTLS(true),
)
```

**Real Examples**: HTTP Client config, Query builders, Test fixtures

**Interview Tip**: Avoids telescoping constructors, makes optional params clear

---

### 4. Prototype
**Problem**: Create new objects by cloning existing ones (when creation is expensive).

**Go Implementation**: Implement `Clone()` method with deep copy.

```go
type Prototype interface {
    Clone() Prototype
}

func (d *Document) Clone() Prototype {
    // Deep copy maps and slices!
    newMetadata := make(map[string]string)
    for k, v := range d.Metadata {
        newMetadata[k] = v
    }
    return &Document{
        Title:    d.Title,
        Metadata: newMetadata,
    }
}
```

**Real Examples**: Document templates, Game object spawning, Config templates

**Interview Tip**: Emphasize deep copy vs shallow copy for slices/maps

---

### 5. Abstract Factory
**Problem**: Create **families** of related objects without specifying concrete classes.

**Go Implementation**: Factory interface with multiple create methods.

```go
type GUIFactory interface {
    CreateButton() Button
    CreateCheckbox() Checkbox
    CreateDialog() Dialog
}

type WindowsFactory struct{}
func (f *WindowsFactory) CreateButton() Button { return &WinButton{} }
func (f *WindowsFactory) CreateCheckbox() Checkbox { return &WinCheckbox{} }

type MacOSFactory struct{}
func (f *MacOSFactory) CreateButton() Button { return &MacButton{} }
func (f *MacOSFactory) CreateCheckbox() Checkbox { return &MacCheckbox{} }
```

**Real Examples**: Cross-platform UI, Database vendors (MySQL/Postgres), Payment providers

**Interview Tip**: Factory creates ONE product type; Abstract Factory creates FAMILIES

---

## Pattern Comparison

```
┌────────────────────┬────────────────────────────────────────────┐
│ Pattern            │ Creates                                    │
├────────────────────┼────────────────────────────────────────────┤
│ Singleton          │ Single instance, globally accessible       │
│ Factory            │ One object, type decided at runtime        │
│ Builder            │ Complex object, step-by-step               │
│ Prototype          │ Copy of existing object                    │
│ Abstract Factory   │ Family of related objects                  │
└────────────────────┴────────────────────────────────────────────┘
```

## Practice Order

```
1. Singleton     → Easiest, understand sync.Once
2. Factory       → Core pattern, interfaces
3. Builder       → Method chaining, optional params
4. Prototype     → Deep copy, Clone interface
5. Abstract Factory → Combines factory concepts
```

## Directory Structure

```
creational/
├── singleton/
│   ├── singleton.go
│   └── singleton_test.go
├── factory/
│   ├── factory.go
│   └── factory_test.go
├── builder/
│   ├── builder.go
│   └── builder_test.go
├── prototype/
│   ├── prototype.go
│   └── prototype_test.go
└── abstract_factory/
    ├── abstract_factory.go      # UI example
    ├── database_factory.go      # Database example
    └── *_test.go
```
