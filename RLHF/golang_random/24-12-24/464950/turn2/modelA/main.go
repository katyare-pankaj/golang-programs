package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Configuration represents the configuration structure.
type Configuration struct {
	AppName string `json:"appName"`
	Port    int    `json:"port"`
}

// ConfigManager represents the configuration manager.
type ConfigManager struct {
	config                    Configuration
	isConfigLoaded            atomic.Bool
	loadConfigWorkerDone      chan struct{}
	ctx                       context.Context
	cancel                    context.CancelFunc
	loadDuration              time.Duration
	loadConfigurationWorkerWG sync.WaitGroup
}

// NewConfigManager creates a new ConfigManager instance with specified load duration.
func NewConfigManager(loadDuration time.Duration) *ConfigManager {
	ctx, cancel := context.WithCancel(context.Background())
	cm := &ConfigManager{
		config:                    Configuration{},
		loadConfigWorkerDone:      make(chan struct{}),
		ctx:                       ctx,
		cancel:                    cancel,
		loadDuration:              loadDuration,
		loadConfigurationWorkerWG: sync.WaitGroup{},
	}
	cm.loadConfigurationAsync()
	return cm
}

// loadConfigurationAsync asynchronously loads the configuration.
func (cm *ConfigManager) loadConfigurationAsync() {
	defer cm.loadConfigurationWorkerWG.Done()

	// Simulate the asynchronous load of configuration.
	time.Sleep(cm.loadDuration)
	config := Configuration{
		AppName: "MyApp",
		Port:    8080,
	}
	cm.setConfiguration(config)
	close(cm.loadConfigWorkerDone)
}

// setConfiguration sets the configuration.
func (cm *ConfigManager) setConfiguration(config Configuration) {
	atomic.StoreBool(&cm.isConfigLoaded, true)
	cm.config = config
}

// loadConfigurationBlocking blocks the caller until the configuration is fully loaded.
func (cm *ConfigManager) loadConfigurationBlocking() {
	select {
	case <-cm.loadConfigWorkerDone:
		// Configuration is already loaded.
		return
	case <-cm.ctx.Done():
		// Context was canceled before load was complete.
		fmt.Println("Configuration loading canceled")
		return
	default:
		// Block until the load is complete.
		<-cm.loadConfigWorkerDone
	}
}

// GetConfiguration gets the loaded configuration.
func (cm *ConfigManager) GetConfiguration() Configuration {
	// Block until configuration is loaded if necessary.
	cm.loadConfigurationBlocking()
	return cm.config
}

// Close shuts down the configuration manager.
func (cm *ConfigManager) Close() {
	cm.cancel()
	cm.loadConfigurationWorkerWG.Wait()
}

func main() {
	cm := NewConfigManager(2 * time.Second)

	// Simulate main application's startup code that needs configurations
	fmt.Println("Starting application...")
	time.Sleep(1 * time.Second)

	// Get configuration immediately and proceed with startup
	config := cm.GetConfiguration()
	if !cm.isConfigLoaded.Load() {
		fmt.Println("Configuration not loaded yet. Aborting startup.")
		cm.Close()
		return
	}

	// Proceed with application logic now that configuration is loaded
	fmt.Println("Configuration loaded:", config)
	fmt.Println("Application running with configuration:", config)
	time.Sleep(5 * time.Second)

	cm.Close()
}
