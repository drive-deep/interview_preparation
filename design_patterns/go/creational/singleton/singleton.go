package singleton

// TODO: Implement Singleton pattern
// 
// Steps:
// 1. Create unexported struct (e.g., logger)
// 2. Create package-level var: instance *logger
// 3. Create sync.Once var: once sync.Once
// 4. Create GetInstance() function using once.Do()
// 5. Add methods to the singleton

import (
	"fmt"
	"sync"
)

type Singleton struct {
	value string 
}

func (s *Singleton) ShowValue() {
	fmt.Println(s.value)
}

var instance *Singleton
var once sync.Once

func GetInstance(value string) *Singleton {
	once.Do(func() {
		instance = &Singleton{value: value}
	})
	return instance
}

// ResetForTesting resets the singleton for testing purposes
// This should only be used in tests!
func ResetForTesting() {
	instance = nil
	once = sync.Once{}
}
