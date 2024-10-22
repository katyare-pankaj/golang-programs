package main

import (
	"fmt"
	"sync"
)

// Singleton struct
type Singleton struct {
	data string
	once sync.Once
}

var singletonInstance *Singleton

func GetInstance() *Singleton {
	singletonInstance.once.Do(func() {
		singletonInstance = &Singleton{}
	})
	return singletonInstance
}

func (s *Singleton) SetData(d string) {
	s.data = d
}

func (s *Singleton) GetData() string {
	return s.data
}

func main() {
	// Usage:
	instance := GetInstance()
	instance.SetData("Hello from Singleton!")
	fmt.Println(instance.GetData()) // Output: Hello from Singleton!
}
