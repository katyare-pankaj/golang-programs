package main

import (
	"fmt"
	"sync"
	"time"
)

// Configuration represents the configuration structure.
type Configuration struct {
	AppName string
	Port    int
}

// AsyncConfigManager represents the configuration manager.
type AsyncConfigManager struct {
	config *Configuration
	wg     sync.WaitGroup
	lock   sync.Mutex
	done   chan struct{}
}

// NewAsyncConfigManager creates a new AsyncConfigManager instance.
func NewAsyncConfigManager() *AsyncConfigManager {
	return &AsyncConfigManager{
		config: &Configuration{},
		wg:     sync.WaitGroup{},
		lock:   sync.Mutex{},
		done:   make(chan struct{}),
	}
}

// LoadConfigAsync initiates the asynchronous configuration loading process.
// It takes a callback function that will be called when the configuration is loaded.
func (m *AsyncConfigManager) LoadConfigAsync(callback func(*Configuration)) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		// Simulate loading configuration from an external source (e.g., file, database)
		// This operation could take some time, so we run it asynchronously.
		time.Sleep(2 * time.Second)
		config := &Configuration{
			AppName: "MyApp",
			Port:    8080,
		}
		m.lock.Lock()
		m.config = config
		m.lock.Unlock()
		// Call the callback function with the loaded configuration
		callback(config)
	}()
}

// WaitForConfigBlocking waits for the configuration to be loaded and returns it.
// It blocks the main goroutine until the configuration is available.
func (m *AsyncConfigManager) WaitForConfigBlocking() *Configuration {
	m.wg.Wait()
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.config
}

// Close shuts down the configuration manager.
func (m *AsyncConfigManager) Close() {
	close(m.done)
}

func main() {
	cm := NewAsyncConfigManager()
	// Start the asynchronous configuration loading
	cm.LoadConfigAsync(func(config *Configuration) {
		fmt.Println("Configuration loaded:", config)
	})
	// Do other work without blocking on configuration loading
	fmt.Println("Doing other work...")
	// Wait for the configuration to be loaded and use it
	config := cm.WaitForConfigBlocking()
	fmt.Println("Using configuration:", config)
	// Close the configuration manager
	cm.Close()
}
