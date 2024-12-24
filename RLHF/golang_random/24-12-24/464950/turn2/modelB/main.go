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
	wg     sync.WaitGroup
	lock   sync.Mutex
	ch     chan struct{}
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		ch: make(chan struct{}),
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
		m.lock.Lock()
		m.config = config
		m.loaded = true
		m.lock.Unlock()
		close(m.ch) // Signal configuration loaded
	}()
}

// WaitForConfig blocks until the configuration is loaded and returns it.
func (m *ConfigManager) WaitForConfig() *Configuration {
	m.wg.Wait()
	m.lock.Lock()
	defer m.lock.Unlock()
	// If the configuration is not loaded, wait for it to be loaded.
	for !m.loaded {
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
	fmt.Println("Doing some work...")

	// Load configuration asynchronously
	cm.LoadConfigAsync()
	fmt.Println("Asynchronous configuration loading initiated.")

	// Other work can be done concurrently while configuration is being loaded.
	time.Sleep(time.Second)
	fmt.Println("Doing some more work...")

	// Later, when needed, the application can block until the configuration is loaded using WaitForConfig.
	config := cm.WaitForConfig()
	fmt.Println("Using configuration:", config)

	// Alternatively, load configuration synchronously
	config = cm.LoadConfig()
	fmt.Println("Using configuration (synchronously):", config)
}
