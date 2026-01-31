# Elevator System - Go

## Requirements
- Multiple elevators
- Handle internal/external requests
- Efficient scheduling
- Display floor & direction

## Patterns
- **State**: Idle, moving up/down, door open
- **Strategy**: Scheduling (SCAN, LOOK, FCFS)
- **Observer**: Display updates
- **Command**: Request queue

## Concurrency Focus
- Elevators as goroutines
- Request queue with channels
- State sync with mutex

## Entities
Building, Elevator, Floor, Request, Scheduler
