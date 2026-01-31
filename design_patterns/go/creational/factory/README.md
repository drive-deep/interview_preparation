# Factory Pattern

## Intent
Define an interface for creating objects, but let subclasses/functions decide which class to instantiate.

---

## 🎤 Interview Questions & Answers

### Q1: What is the Factory Pattern?
**A:** Factory Pattern is a creational design pattern that provides an interface for creating objects without specifying their concrete classes. Instead of using `new` directly, we delegate object creation to a factory function that returns objects based on input parameters.

### Q2: Why do we use Factory Pattern?
**A:** We use Factory Pattern when:
- We don't know beforehand which concrete class we need
- We want to **decouple** object creation from usage
- We need to **centralize** creation logic in one place
- We want to easily **add new types** without changing client code

**Example:** A notification system that can send via Email, SMS, or Push. The client just calls `NotifierFactory("email")` without knowing how `EmailNotifier` is created.

### Q3: What problem does Factory Pattern solve?
**A:** It solves the problem of **tight coupling** between client code and concrete implementations.

**Without Factory (Bad):**
```go
// Client is tightly coupled to EmailNotifier
notifier := &EmailNotifier{Vendor: "gmail"}
```

**With Factory (Good):**
```go
// Client depends only on Notifier interface
notifier, _ := NotifierFactory("email", "gmail")
```

### Q4: When should I NOT use Factory Pattern?
**A:** Don't use it when:
- Only one implementation exists
- Creation logic is trivial (`&Struct{}` is enough)
- You don't need polymorphism
- It adds unnecessary complexity

### Q5: Factory vs Abstract Factory - What's the difference?
**A:** 
| Factory | Abstract Factory |
|---------|------------------|
| Creates ONE product | Creates FAMILY of related products |
| Single factory function | Factory interface with multiple implementations |
| `NotifierFactory("email")` | `UIFactory.CreateButton()`, `UIFactory.CreateDialog()` |

### Q6: How is Factory Pattern implemented in Go differently than Java?
**A:**
| Java | Go |
|------|-----|
| Abstract Factory class | Simple function |
| Inheritance hierarchy | Interface + composition |
| `new ConcreteProduct()` | `&ConcreteProduct{}` |
| Factory Method pattern | Just return interface from function |

Go doesn't need complex class hierarchies - a simple function returning an interface is enough.

### Q7: Give a real-world example of Factory Pattern.
**A:** 
- **`database/sql`**: `sql.Open("mysql", connStr)` - factory creates different DB drivers
- **`http`**: `http.NewRequest("GET", url, nil)` - creates different request types
- **Payment gateways**: `PaymentFactory("stripe")` returns Stripe/PayPal/Razorpay processor

---

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
    │  notifier.Send("hello")       │     (polymorphic call)         │
```

---

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

---

## Advantages & Disadvantages

| ✅ Advantages | ❌ Disadvantages |
|--------------|-----------------|
| Loose coupling | Can lead to many classes |
| Single Responsibility | Factory can grow large |
| Open/Closed Principle | Slight indirection overhead |
| Easy to test/mock | |

---

## Real-World Examples

| Example | Factory Creates |
|---------|-----------------|
| `sql.Open("mysql", ...)` | MySQL, PostgreSQL, SQLite |
| Notification services | Email, SMS, Push notifiers |
| Payment processors | Stripe, PayPal, Razorpay |
| Loggers | File, Console, Remote loggers |

---

## Files
- `factory.go` - Implementation
- `factory_test.go` - Tests (positive + negative)
