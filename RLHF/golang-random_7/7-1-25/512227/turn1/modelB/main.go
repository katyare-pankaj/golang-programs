package main

import (
	"fmt"
	"reflect"
	"time"
)

type WeatherData struct {
	City        string
	Temperature float32
	Humidity    float32
	Pressure    float32
	Timestamp   time.Time
}

func main() {
	weatherData := []WeatherData{
		{City: "New York", Temperature: 25.6, Humidity: 68.4, Pressure: 1012.0, Timestamp: time.Now()},
		{City: "London", Temperature: 18.2, Humidity: 79.0, Pressure: 1010.0, Timestamp: time.Now().Add(-time.Hour)},
		{City: "Paris", Temperature: 22.0, Humidity: 65.3, Pressure: 1013.0, Timestamp: time.Now().Add(-2 * time.Hour)},
	}

	//Example of updating data dynamically using reflection
	updateWeatherData(weatherData, "London", 20.0, 70.0, 1011.0)

	//Example of generating a report dynamically using reflection
	generateWeatherReport(weatherData)
}

// Function to update weather data dynamically using reflection
func updateWeatherData(data interface{}, city string, temperature float32, humidity float32, pressure float32) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		panic("data must be a slice")
	}

	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		fmt.Println("Elem: ", elem)
		cityField := elem.FieldByName("City")
		if cityField.Kind() == reflect.String && cityField.String() == city {
			temperatureField := elem.FieldByName("Temperature")
			temperatureField.SetFloat(float64(temperature))

			humidityField := elem.FieldByName("Humidity")
			humidityField.SetFloat(float64(humidity))

			pressureField := elem.FieldByName("Pressure")
			pressureField.SetFloat(float64(pressure))

			timestampField := elem.FieldByName("Timestamp")
			timestampField.Set(reflect.ValueOf(time.Now()))

			fmt.Println("Weather data updated for", city)
			return
		}
	}
	fmt.Println("City not found in weather data.")
}

// Function to generate a weather report dynamically using reflection
func generateWeatherReport(data interface{}) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		panic("data must be a slice")
	}

	fmt.Println("Weather Report:")
	fmt.Println("----------------")
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		city := elem.FieldByName("City").String()
		temperature := elem.FieldByName("Temperature").Float()
		humidity := elem.FieldByName("Humidity").Float()
		pressure := elem.FieldByName("Pressure").Float()
		timestamp := elem.FieldByName("Timestamp").Interface().(time.Time)

		fmt.Printf("City: %s\n", city)
		fmt.Printf("Temperature: %.1f°C\n", temperature)
		fmt.Printf("Humidity: %.1f%%\n", humidity)
		fmt.Printf("Pressure: %.1f hPa\n", pressure)
		fmt.Printf("Timestamp: %s\n\n", timestamp.Format(time.RFC3339))
	}
}
