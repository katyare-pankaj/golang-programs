package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Features Features `json:"features"`
}

type Features struct {
	NewFeature bool `json:"new_feature"`
	DebugMode  bool `json:"debug_mode"`
}

// LoadConfig loads the configuration from a file.
func LoadConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
}
