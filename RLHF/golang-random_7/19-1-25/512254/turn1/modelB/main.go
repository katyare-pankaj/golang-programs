package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// PrivacyLevel represents the data privacy level of a query parameter.
type PrivacyLevel int

const (
	// Private indicates that the parameter should be encrypted and only accessible to authorized parties.
	Private PrivacyLevel = iota
	// Confidential indicates that the parameter should be encrypted and only accessible to authorized parties with specific roles.
	Confidential
	// Public indicates that the parameter can be stored and shared with third parties.
	Public
)

// ParamConfig defines the configuration for a query parameter.
type ParamConfig struct {
	PrivacyLevel PrivacyLevel `json:"privacyLevel"`
	Description  string       `json:"description"`
}

// QueryParamManager manages URL query parameters with privacy, security, and compliance strategies.
type QueryParamManager struct {
	config map[string]ParamConfig
}

// NewQueryParamManager creates a new QueryParamManager with the given configuration.
func NewQueryParamManager(config map[string]ParamConfig) *QueryParamManager {
	return &QueryParamManager{config: config}
}

// GetPrivacyLevel retrieves the privacy level of a given query parameter.
func (m *QueryParamManager) GetPrivacyLevel(param string) PrivacyLevel {
	config, ok := m.config[param]
	if !ok {
		return Private // Default privacy level for unrecognized parameters
	}
	return config.PrivacyLevel
}

// EncryptQuery encrypts the values of query parameters based on their privacy levels.
func (m *QueryParamManager) EncryptQuery(query url.Values) url.Values {
	encryptedQuery := url.Values{}

	for param, values := range query {
		privacyLevel := m.GetPrivacyLevel(param)
		for _, value := range values {
			switch privacyLevel {
			case Private, Confidential:
				// Encrypt the value here
				encryptedValue := encrypt(value)
				encryptedQuery.Add(param, encryptedValue)
			default:
				encryptedQuery.Add(param, value)
			}
		}
	}

	return encryptedQuery
}

// DecryptQuery decrypts the values of query parameters based on their privacy levels.
func (m *QueryParamManager) DecryptQuery(query url.Values) url.Values {
	decryptedQuery := url.Values{}

	for param, values := range query {
		privacyLevel := m.GetPrivacyLevel(param)
		for _, value := range values {
			switch privacyLevel {
			case Private, Confidential:
				// Decrypt the value here
				decryptedValue := decrypt(value)
				decryptedQuery.Add(param, decryptedValue)
			default:
				decryptedQuery.Add(param, value)
			}
		}
	}

	return decryptedQuery
}

func encrypt(value string) string {
	// Placeholder encryption logic. In a real implementation, use a secure encryption algorithm.
	return strings.ReplaceAll(value, "a", "z")
}

func decrypt(value string) string {
	// Placeholder decryption logic. In a real implementation, use the same encryption algorithm used for encryption.
	return strings.ReplaceAll(value, "z", "a")
}

func main() {
	// Example query parameters and their configurations
	config := map[string]ParamConfig{
		"customer_id":  {PrivacyLevel: Private, Description: "Customer ID for personalized content"},
		"email":        {PrivacyLevel: Confidential, Description: "Customer email for marketing purposes"},
		"location":     {PrivacyLevel: Public, Description: "Customer location for analytics"},
		"app_version": {PrivacyLevel: Public, Description: "Application version for troubleshooting"},
	}

	qpm := NewQueryParamManager(config)

	// Example query with sensitive parameters
	originalQuery := url.Values{}
	originalQuery.Add("customer_id", "12345")
	originalQuery.Add("email", "customer@example.com")