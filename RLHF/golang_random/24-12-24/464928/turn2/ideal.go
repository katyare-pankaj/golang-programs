package main

import (
	"fmt"
	"sync"
	"time"
)

type Configuration struct {
	AppName string
	Port    int
}

// ConfigManager represents the configuration manager.
type ConfigManager struct {
	config *Configuration
	loaded bool
	mu     sync.Mutex
	wg     sync.WaitGroup
	ch     chan struct{}
}

// NewConfigManager initializes a new ConfigManager.
func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		ch: make(chan struct{}, 1), // Buffered channel to avoid blocking
	}
}

// LoadConfigAsync initiates the asynchronous configuration loading process.
func (m *ConfigManager) LoadConfigAsync() {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		time.Sleep(2 * time.Second) // Simulate loading time
		config := &Configuration{
			AppName: "MyApp",
			Port:    8080,
		}

		m.mu.Lock()
		m.config = config
		m.loaded = true
		m.mu.Unlock()

		// Notify that the configuration is loaded.
		select {
		case m.ch <- struct{}{}:
		default:
		}
	}()
}

// WaitForConfig blocks until the configuration is loaded and returns it.
func (m *ConfigManager) WaitForConfig() *Configuration {
	m.wg.Wait()
	m.mu.Lock()
	defer m.mu.Unlock()

	// Wait for the configuration to be loaded if not already
	if !m.loaded {
		<-m.ch
	}

	return m.config
}

// LoadConfig loads the configuration synchronously.
func (m *ConfigManager) LoadConfig() *Configuration {
	m.LoadConfigAsync()
	return m.WaitForConfig()
}

func main() {
	cm := NewConfigManager()

	fmt.Println("Starting asynchronous configuration loading...")

	// Load configuration asynchronously
	cm.LoadConfigAsync()

	// Simulate doing some other work
	time.Sleep(1 * time.Second)
	fmt.Println("Doing some other work...")

	// Later, block and wait for the configuration to be ready
	config := cm.WaitForConfig()
	fmt.Println("Configuration loaded:", config)

	// Alternatively, load configuration synchronously
	config = cm.LoadConfig()
	fmt.Println("Configuration loaded synchronously:", config)
}
