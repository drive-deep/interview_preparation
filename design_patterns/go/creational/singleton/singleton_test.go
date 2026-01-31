package singleton

import (
	"fmt"
	"sync"
	"testing"
)

// TODO: Write tests for Singleton pattern
//
// Test cases to implement:
// 1. TestGetInstance_ReturnsSameInstance
// 2. TestGetInstance_ConcurrentAccess (use goroutines + WaitGroup)
// 3. TestGetInstance_NotNil
// 4. TestSingleton_StatePersists
func TestGetInstance_ReturnsSameInstance(t *testing.T) {
	instance1 := GetInstance("first")
	instance2 := GetInstance("second")

	if instance1 != instance2 {
		t.Errorf("Expected both instances to be the same, but they are different")
	}
}

func TestGetInstance_NotNil(t *testing.T) {
	instance := GetInstance("test")

	if instance == nil {
		t.Errorf("Expected instance to be not nil")
	}
}

func TestSingleton_StatePersists(t *testing.T) {
	instance1 := GetInstance("first")
	instance2 := GetInstance("second")

	if instance1.value != "first" {
		t.Errorf("Expected instance value to be 'first', got '%s'", instance1.value)
	}

	if instance2.value != "first" {
		t.Errorf("Expected instance value to be 'first', got '%s'", instance2.value)
	}
}

func TestGetInstance_ConcurrentAccess(t *testing.T) {
	// IMPORTANT: Reset singleton to test concurrent initialization
	// This resets both instance AND sync.Once
	ResetForTesting()

	const goroutineCount = 1000 // More goroutines = higher chance of race
	results := make(chan *Singleton, goroutineCount)

	// Use WaitGroup to make all goroutines start at the same time
	var wg sync.WaitGroup
	startSignal := make(chan struct{})

	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-startSignal // Wait for signal to start simultaneously
			inst := GetInstance(fmt.Sprintf("goroutine-%d", id))
			results <- inst
		}(i)
	}

	// Release all goroutines at once - this maximizes race condition chance
	close(startSignal)

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	var firstInstance *Singleton
	var firstValue string
	for inst := range results {
		if firstInstance == nil {
			firstInstance = inst
			firstValue = inst.value
		} else if firstInstance != inst {
			t.Errorf("RACE DETECTED: Got different instance pointers!")
		}
	}

	// Also check if value was overwritten (another sign of race)
	if firstInstance != nil && firstInstance.value != firstValue {
		t.Errorf("RACE DETECTED: Value changed from '%s' to '%s'", firstValue, firstInstance.value)
	}
}
