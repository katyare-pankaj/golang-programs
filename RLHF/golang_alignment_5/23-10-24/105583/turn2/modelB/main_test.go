package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// WeatherService is an interface for accessing the weather API.
type WeatherService interface {
	GetWeather(ctx context.Context, city string) (string, error)
}

// MockWeatherService is a mock implementation of the WeatherService interface.
type MockWeatherService struct {
	mock.Mock
}

// GetWeather is a mock method for retrieving weather data.
func (m *MockWeatherService) GetWeather(ctx context.Context, city string) (string, error) {
	args := m.Called(ctx, city)
	return args.String(0), args.Error(1)
}

// DistributedSystem is the main struct that handles the integration with the third-party service.
type DistributedSystem struct {
	weatherService WeatherService
}

// NewDistributedSystem creates a new instance of the DistributedSystem.
func NewDistributedSystem(weatherService WeatherService) *DistributedSystem {
	return &DistributedSystem{weatherService: weatherService}
}

// GetWeatherForCity is a method that uses the WeatherService to retrieve weather data for a given city.
func (ds *DistributedSystem) GetWeatherForCity(ctx context.Context, city string) (string, error) {
	return ds.weatherService.GetWeather(ctx, city)
}

func TestGetWeatherForCity(t *testing.T) {
	t.Run("Successful Weather Retrieval", func(t *testing.T) {
		// Set up the mock weather service
		mockWeatherService := &MockWeatherService{}
		mockWeatherService.On("GetWeather", mock.Anything, "London").Return("Cloudy", nil)

		// Create the distributed system with the mock weather service
		ds := NewDistributedSystem(mockWeatherService)

		// Call the method under test
		weather, err := ds.GetWeatherForCity(context.Background(), "London")

		// Verify the results
		assert.NoError(t, err)
		assert.Equal(t, "Cloudy", weather)
		mockWeatherService.AssertExpectations(t)
	})

	t.Run("Weather API Error", func(t *testing.T) {
		// Set up the mock weather service with an error
		mockWeatherService := &MockWeatherService{}
		mockWeatherService.On("GetWeather", mock.Anything, "Paris").Return("", fmt.Errorf("API error"))

		// Create the distributed system with the mock weather service
		ds := NewDistributedSystem(mockWeatherService)

		// Call the method under test
		weather, err := ds.GetWeatherForCity(context.Background(), "Paris")

		// Verify the results
		assert.Error(t, err)
		assert.Equal(t, "", weather)
		mockWeatherService.AssertExpectations(t)
	})
}

func Example() {
	// In a real application, you would use an actual implementation of the WeatherService
	// that interacts with the third-party weather API.
	weatherService := &RealWeatherService{}
	ds := NewDistributedSystem(weatherService)

	// Call the method under test with a real city
	weather, err := ds.GetWeatherForCity(context.Background(), "New York")
	if err != nil {
		log.Fatalf("Error getting weather: %v", err)
	}
	fmt.Printf("Weather in New York: %s\n", weather)
}

type RealWeatherService struct{}

func (s *RealWeatherService) GetWeather(ctx context.Context, city string) (string, error) {
	// Implement the actual integration with the third-party weather API
	// For simplicity, we'll just return a dummy weather data in this example.
	return "Sunny", nil
}
