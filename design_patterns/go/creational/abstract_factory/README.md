# Abstract Factory Pattern

## Intent
Provide an interface for creating **families of related objects** without specifying their concrete classes.

## Key Difference from Factory
| Factory | Abstract Factory |
|---------|------------------|
| Creates **one type** of object | Creates **families** of related objects |
| Single factory method | Multiple factory methods per factory |
| `NotifierFactory("email")` → EmailNotifier | `UIFactory.CreateButton()` + `UIFactory.CreateCheckbox()` |

## When to Use
- System needs to be independent of how its products are created
- System should work with **multiple families** of products
- Related products are designed to be used together
- You want to provide a library of products, revealing only interfaces

## Real-World Examples
- **Cross-platform UI**: Windows vs macOS buttons, checkboxes, dialogs
- **Database drivers**: MySQL vs PostgreSQL connections, queries, transactions
- **Document export**: PDF vs Word vs HTML (header, body, footer components)
- **Theme systems**: Dark vs Light theme components
- **Payment gateways**: Stripe vs Razorpay vs PayPal

## Example: Payment Gateway

```go
// Product Interfaces
type Payment interface {
    Process(amount float64) error
}
type Refund interface {
    Process(transactionID string) error
}

// Abstract Factory
type PaymentFactory interface {
    CreatePayment() Payment
    CreateRefund() Refund
}

// Stripe Family
type StripePayment struct{}
func (s *StripePayment) Process(amount float64) error {
    // Stripe API: POST /v1/charges
    return nil
}

type StripeRefund struct{}
func (s *StripeRefund) Process(txnID string) error {
    // Stripe API: POST /v1/refunds
    return nil
}

type StripeFactory struct{}
func (f *StripeFactory) CreatePayment() Payment { return &StripePayment{} }
func (f *StripeFactory) CreateRefund() Refund   { return &StripeRefund{} }

// Razorpay Family
type RazorpayPayment struct{}
func (r *RazorpayPayment) Process(amount float64) error {
    // Razorpay API: POST /v1/orders
    return nil
}

type RazorpayRefund struct{}
func (r *RazorpayRefund) Process(txnID string) error {
    // Razorpay API: POST /v1/payments/{id}/refund
    return nil
}

type RazorpayFactory struct{}
func (f *RazorpayFactory) CreatePayment() Payment { return &RazorpayPayment{} }
func (f *RazorpayFactory) CreateRefund() Refund   { return &RazorpayRefund{} }

// Usage - swap gateway without changing business logic
func ProcessOrder(factory PaymentFactory, amount float64) {
    payment := factory.CreatePayment()
    payment.Process(amount)
}
```

**Benefits**: Switch from Stripe to Razorpay by just changing the factory - zero changes to business logic!

## Structure
```
┌─────────────────────────┐
│   AbstractFactory       │ ← Interface
│  + CreateButton()       │
│  + CreateCheckbox()     │
└───────────┬─────────────┘
            │
    ┌───────┴───────┐
    │               │
┌───▼───┐       ┌───▼───┐
│Windows│       │ macOS │
│Factory│       │Factory│
└───┬───┘       └───┬───┘
    │               │
    ▼               ▼
WinButton       MacButton
WinCheckbox     MacCheckbox
```

## Go Implementation Notes
- Use interfaces for both factory and products
- Factory returns product interfaces, not concrete types
- Client code works only with abstractions

## Interview Tips
1. **vs Factory Method**: Abstract Factory creates families; Factory creates one type
2. **vs Builder**: Builder constructs complex objects step-by-step; Abstract Factory creates related objects at once
3. **Open/Closed Principle**: Adding new family = new factory (no existing code changes)
4. **Single Responsibility**: Each factory handles one product family
