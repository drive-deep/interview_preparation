# Prototype Pattern

## Intent
Create new objects by **cloning existing objects** instead of creating from scratch.

## When to Use
- Object creation is expensive (DB calls, file reads, complex computation)
- You need many similar objects with slight variations
- You want to avoid subclasses of an object creator
- Runtime object configuration is preferred over compile-time

## Real-World Examples
- **Document templates**: Clone a base Word/PDF template, then customize
- **Game objects**: Clone enemy/NPC prototypes instead of recreating
- **Database records**: Clone a config record, modify specific fields
- **UI components**: Clone a styled component, change just the text

## Structure
```
┌──────────────────┐
│   Prototype      │ ← Interface
│  + Clone()       │
└────────┬─────────┘
         │
    ┌────┴────┐
    ▼         ▼
┌────────┐ ┌────────┐
│Concrete│ │Concrete│
│Proto A │ │Proto B │
└────────┘ └────────┘
```

## Go Implementation Notes
```go
// Prototype interface
type Prototype interface {
    Clone() Prototype
}

// Concrete prototype
type Document struct {
    Title   string
    Content string
    Author  string
}

func (d *Document) Clone() Prototype {
    // Deep copy - create new object with same values
    return &Document{
        Title:   d.Title,
        Content: d.Content,
        Author:  d.Author,
    }
}

// Usage
baseTemplate := &Document{Title: "Report", Content: "...", Author: "System"}
myDoc := baseTemplate.Clone().(*Document)
myDoc.Author = "Prabhat"  // Customize the clone
```

## Deep vs Shallow Copy
| Type | Description | When to Use |
|------|-------------|-------------|
| **Shallow** | Copies pointers, shared references | Simple structs, immutable fields |
| **Deep** | Copies all nested objects recursively | Complex objects with slices/maps/pointers |

## Interview Tips
1. **vs Factory**: Factory creates from scratch; Prototype clones existing
2. **Performance**: Useful when object creation is expensive
3. **Go idiom**: Implement `Clone()` method that returns interface type
4. **Careful with pointers**: Always deep copy slices, maps, and pointer fields

---

## TODO: Implement

- [ ] `prototype.go` - Implement Document/Config prototype with Clone()
- [ ] `prototype_test.go` - Test cloning, verify deep copy works
- [ ] Add a second example (e.g., game character prototype)
