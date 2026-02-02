# Structural Design Patterns in Go

## What Are Structural Patterns?

Structural patterns deal with **HOW objects are composed and connected** to form larger, flexible structures. They ensure that when one part changes, the entire structure doesn't need to change.

```
┌─────────────────────────────────────────────────────────────────┐
│                    STRUCTURAL PATTERNS                          │
│                                                                 │
│   "Compose objects into larger structures while keeping them    │
│    flexible and efficient"                                      │
│                                                                 │
│   Creational = HOW to create                                    │
│   Structural = HOW to compose  ← You are here                   │
│   Behavioral = HOW to communicate                               │
└─────────────────────────────────────────────────────────────────┘
```

## When to Use Structural Patterns

| Situation | Pattern to Use |
|-----------|----------------|
| Make incompatible interfaces work together | **Adapter** |
| Add behavior dynamically without modifying | **Decorator** |
| Simplify complex subsystem | **Facade** |
| Control access to an object | **Proxy** |
| Treat individual and groups uniformly | **Composite** |
| Separate abstraction from implementation | **Bridge** |

---

## The 6 Structural Patterns

### 1. Adapter (Wrapper)
**Problem**: Make incompatible interfaces work together.

**Go Implementation**: Wrapper struct that implements target interface.

```go
// Target interface your code expects
type PaymentProcessor interface {
    ProcessPayment(amount float64) error
}

// Third-party Stripe SDK (incompatible interface)
type StripeSDK struct{}
func (s *StripeSDK) Charge(cents int, currency string) error { ... }

// Adapter makes Stripe work with your interface
type StripeAdapter struct {
    stripe *StripeSDK
}

func (a *StripeAdapter) ProcessPayment(amount float64) error {
    cents := int(amount * 100)
    return a.stripe.Charge(cents, "USD")
}
```

**Real Examples**: Legacy API integration, Third-party SDK wrappers, Database driver adapters

**Interview Tip**: "Convert interface A to interface B without modifying either"

---

### 2. Decorator (Middleware)
**Problem**: Add behavior to objects dynamically without inheritance.

**Go Implementation**: Middleware pattern - wrap handlers with additional functionality.

```go
type Handler interface {
    Handle(request string) string
}

// Core handler
type CoreHandler struct{}
func (h *CoreHandler) Handle(request string) string {
    return "Response: " + request
}

// Logging decorator
type LoggingDecorator struct {
    wrapped Handler
}
func (d *LoggingDecorator) Handle(request string) string {
    log.Println("Request:", request)
    result := d.wrapped.Handle(request)
    log.Println("Response:", result)
    return result
}

// Usage: chain decorators
handler := &LoggingDecorator{
    wrapped: &AuthDecorator{
        wrapped: &CoreHandler{},
    },
}
```

**Real Examples**: HTTP middleware (logging, auth, CORS), io.Reader wrappers, Compression layers

**Interview Tip**: "Same interface, just adding behavior on top"

---

### 3. Facade
**Problem**: Provide a simplified interface to a complex subsystem.

**Go Implementation**: Single struct/package exposing high-level methods.

```go
// Complex subsystems
type InventoryService struct{}
type PaymentService struct{}
type ShippingService struct{}

// Facade - simplified interface
type OrderFacade struct {
    inventory *InventoryService
    payment   *PaymentService
    shipping  *ShippingService
}

func (f *OrderFacade) PlaceOrder(productID string, quantity int) error {
    // Orchestrates all complex subsystems
    if err := f.inventory.Reserve(productID, quantity); err != nil {
        return err
    }
    if err := f.payment.Charge(productID, quantity); err != nil {
        f.inventory.Release(productID, quantity)
        return err
    }
    return f.shipping.Schedule(productID, quantity)
}
```

**Real Examples**: API gateway, SDK wrappers, System initialization

**Interview Tip**: "One door to many rooms"

---

### 4. Proxy
**Problem**: Control access to another object.

**Go Implementation**: Wrapper with same interface that adds access control.

```go
type Database interface {
    Query(sql string) []Row
}

// Real implementation
type RealDatabase struct{}
func (d *RealDatabase) Query(sql string) []Row { ... }

// Caching Proxy
type CachingProxy struct {
    db    Database
    cache map[string][]Row
}

func (p *CachingProxy) Query(sql string) []Row {
    if cached, ok := p.cache[sql]; ok {
        return cached  // Return cached result
    }
    result := p.db.Query(sql)
    p.cache[sql] = result
    return result
}
```

**Proxy Types**:
- **Caching Proxy**: Cache expensive operations
- **Auth Proxy**: Check permissions before access
- **Lazy Proxy**: Delay creation until needed
- **Logging Proxy**: Log all accesses

**Real Examples**: Redis cache proxy, API rate limiter, Lazy image loading

**Interview Tip**: "Same interface as real object, but controls access"

---

### 5. Composite
**Problem**: Treat individual objects and compositions uniformly.

**Go Implementation**: Recursive interface for tree structures.

```go
type FileSystemNode interface {
    GetSize() int64
    GetName() string
}

// Leaf - individual file
type File struct {
    name string
    size int64
}
func (f *File) GetSize() int64 { return f.size }
func (f *File) GetName() string { return f.name }

// Composite - folder containing files/folders
type Folder struct {
    name     string
    children []FileSystemNode
}
func (f *Folder) GetSize() int64 {
    total := int64(0)
    for _, child := range f.children {
        total += child.GetSize()  // Works for both files and folders!
    }
    return total
}
func (f *Folder) GetName() string { return f.name }
```

**Real Examples**: File systems, UI component trees, Organization charts, Menu structures

**Interview Tip**: "Leaf and composite share the same interface"

---

### 6. Bridge
**Problem**: Separate abstraction from implementation so they can vary independently.

**Go Implementation**: Composition with interface for implementation.

```go
// Implementation interface
type Renderer interface {
    RenderCircle(radius float64)
    RenderSquare(side float64)
}

type VectorRenderer struct{}
func (v *VectorRenderer) RenderCircle(r float64) { fmt.Println("Vector circle") }

type RasterRenderer struct{}
func (r *RasterRenderer) RenderCircle(radius float64) { fmt.Println("Raster circle") }

// Abstraction
type Shape struct {
    renderer Renderer  // Bridge to implementation
}

type Circle struct {
    Shape
    radius float64
}
func (c *Circle) Draw() {
    c.renderer.RenderCircle(c.radius)
}

// Usage - can combine any shape with any renderer
circle := &Circle{Shape{renderer: &VectorRenderer{}}, 5.0}
circle.Draw()
```

**Real Examples**: Cross-platform apps, Multiple database backends, Remote controls + devices

**Interview Tip**: "2D matrix: abstractions × implementations"

---

## Pattern Comparison

```
┌────────────────┬────────────────────────────────────────────────┐
│ Pattern        │ Purpose                                        │
├────────────────┼────────────────────────────────────────────────┤
│ Adapter        │ Convert interface A → interface B              │
│ Decorator      │ Add behavior dynamically (same interface)      │
│ Facade         │ Simplify complex subsystem                     │
│ Proxy          │ Control access (same interface)                │
│ Composite      │ Tree structure, uniform treatment              │
│ Bridge         │ Decouple abstraction from implementation       │
└────────────────┴────────────────────────────────────────────────┘
```

## Visual Summary

```
ADAPTER             DECORATOR           FACADE
┌───┐    ┌───┐     ┌───┐→┌───┐→┌───┐   ┌──────┐
│Old│ → │New│     │Log│ │Auth│ │Core│  │Simple│→[Complex]
└───┘    └───┘     └───┘ └───┘ └───┘   │ API  │ [Systems]
"Make it fit"      "Wrap & extend"     └──────┘
                                       "One door"

PROXY               COMPOSITE           BRIDGE
Client→Proxy→Real   ┌─Folder──┐         Abstraction
"Control access"    │File File│          ↓      ↓
                    │ Folder  │        ImplA  ImplB
                    └─────────┘        "Decouple"
                   "Tree uniform"
```

## Practice Order

```
1. Adapter   → Most common, wrap third-party APIs
2. Decorator → Go middleware pattern, very idiomatic
3. Facade    → Simplify complex systems
4. Proxy     → Caching, auth, lazy loading
5. Composite → Tree structures
6. Bridge    → Less common, advanced decoupling
```

## Directory Structure

```
structural/
├── adapter/
│   ├── adapter.go
│   └── adapter_test.go
├── decorator/
│   ├── decorator.go
│   └── decorator_test.go
├── facade/
│   ├── facade.go
│   └── facade_test.go
├── proxy/
│   ├── proxy.go
│   └── proxy_test.go
├── composite/
│   ├── composite.go
│   └── composite_test.go
└── bridge/
    ├── bridge.go
    └── bridge_test.go
```

## Interview Frequency

| Pattern | Interview Frequency | Notes |
|---------|---------------------|-------|
| Adapter | ⭐⭐⭐⭐⭐ | Asked very often |
| Decorator | ⭐⭐⭐⭐⭐ | Go middleware is decorator |
| Facade | ⭐⭐⭐⭐ | Common in system design |
| Proxy | ⭐⭐⭐⭐ | Caching questions |
| Composite | ⭐⭐⭐ | File system questions |
| Bridge | ⭐⭐ | Less common |
