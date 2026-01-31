# Factory Pattern

## Intent
Define an interface for creating objects, but let subclasses/functions decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses.

## Problem It Solves
- Need to create objects without specifying exact concrete class
- Want to decouple object creation from usage
- Need flexibility to add new types without changing client code
- Hide complex creation logic from clients

## How It Works

```
Client Code                      Factory                         Products
    │                               │                                │
    │  NotifierFactory("email")     │                                │
    │ ─────────────────────────────>│                                │
    │                               │  switch "email" → create       │
    │                               │ ───────────────────────────────>│
    │                               │  return &EmailNotifier{}       │
    │                               │<───────────────────────────────│
    │  returns Notifier interface   │                                │
    │<─────────────────────────────│                                │
    │                               │                                │
    │  notifier.Send("hello")       │                                │
    │  (polymorphic call)           │                                │
```

## Real-World Examples
| Example | Factory Creates |
|---------|-----------------|
| Database drivers | `sql.Open("mysql", ...)` → MySQL, PostgreSQL, SQLite |
| Notification services | Email, SMS, Push notifiers |
| Payment processors | Stripe, PayPal, Razorpay handlers |
| Loggers | File, Console, Remote loggers |
| Serializers | JSON, XML, YAML encoders |

## Go-Idiomatic Approach

| Java/OOP Way | Go Way |
|--------------|--------|
| Abstract Factory class | Factory function returning interface |
| Concrete factory subclasses | Switch/map in factory function |
| new ConcreteProduct() | Return &ConcreteProduct{} |
| Factory interface | Simple function (no interface needed) |

## Structure

```
┌─────────────────────────────────┐
│      Notifier (interface)       │
├─────────────────────────────────┤
│ + Send(message string) error    │
└─────────────────────────────────┘
              ▲
              │ implements
       ┌──────┼──────┐
       │      │      │
┌──────────┐ ┌──────────┐ ┌──────────┐
│EmailNoti │ │ SmsNoti  │ │ PushNoti │
└──────────┘ └──────────┘ └──────────┘

┌─────────────────────────────────────────────────┐
│ NotifierFactory(type, vendor) (Notifier, error) │
└─────────────────────────────────────────────────┘
```

## Key Components

| Component | Purpose |
|-----------|---------|
| Product interface | Common interface all products implement (`Notifier`) |
| Concrete Products | Different implementations (`EmailNotifier`, `SmsNotifier`) |
| Factory Function | Creates correct product based on input type |

## Advantages
- ✅ **Loose coupling** - Client depends on interface, not concrete types
- ✅ **Single Responsibility** - Creation logic in one place
- ✅ **Open/Closed Principle** - Add new types without modifying existing code
- ✅ **Testability** - Easy to mock by providing different implementations

## Disadvantages
- ❌ Can lead to many small classes/structs
- ❌ Factory function can become large with many types (use map or registry)

## When to Use
- ✅ Multiple implementations of same interface
- ✅ Object creation logic is complex
- ✅ Need to add new types without changing clients
- ✅ Want to hide concrete types from clients

## When NOT to Use
- ❌ Only one implementation exists
- ❌ Simple object creation (just use `&Struct{}`)
- ❌ No need for polymorphism

## Variations

### Simple Factory (what we implemented)
```go
func NotifierFactory(type string) (Notifier, error)
```

### Factory with Registry (for many types)
```go
var registry = map[string]func() Notifier{
    "email": func() Notifier { return &EmailNotifier{} },
    "sms":   func() Notifier { return &SmsNotifier{} },
}

func NotifierFactory(type string) (Notifier, error) {
    if factory, ok := registry[type]; ok {
        return factory(), nil
    }
    return nil, fmt.Errorf("unknown type: %s", type)
}
```

## Test Cases
### Positive Tests
1. Factory returns correct type for each valid input
2. Returned notifier has correct vendor value
3. Each notifier implements interface correctly
4. Polymorphic behavior works

### Negative Tests
5. Factory returns error for unknown type
6. Factory returns error for empty type
7. Factory handles case sensitivity (EMAIL ≠ email)
8. Factory rejects special characters
9. Factory handles whitespace correctly

## Files
- `factory.go` - Implementation
- `factory_test.go` - Tests (positive + negative)
