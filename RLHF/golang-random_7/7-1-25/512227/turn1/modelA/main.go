package weather_app

import (
	"fmt"
	"reflect"
)

// WeatherData is a base interface to hold various weather information.
type WeatherData interface {
	Description() string
}

// TemperatureData represents temperature data.
type TemperatureData struct {
	Degrees float32 `json:"degrees"`
	Unit    string  `json:"unit"`
}

func (t *TemperatureData) Description() string {
	return fmt.Sprintf("Temperature: %.2f %s", t.Degrees, t.Unit)
}

// HumidityData represents humidity data.
type HumidityData struct {
	Value float32 `json:"value"`
}

func (h *HumidityData) Description() string {
	return fmt.Sprintf("Humidity: %.2f%%", h.Value)
}

// ProcessWeatherData uses reflection to process and display weather data dynamically.
func ProcessWeatherData(data []WeatherData) {
	for _, entry := range data {
		value := reflect.ValueOf(entry)
		typeOf := value.Type()

		for i := 0; i < typeOf.NumField(); i++ {
			field := typeOf.Field(i)
			tag := field.Tag
			key := tag.Get("json")
			if key != "" {
				fmt.Printf("%s: %s\n", key, value.Field(i).String())
			}
		}

		fmt.Println(entry.Description())
		fmt.Println("-----")
	}
}

func main() {
	weatherData := []WeatherData{
		&TemperatureData{Degrees: 25.0, Unit: "°C"},
		&HumidityData{Value: 60.0},
	}

	ProcessWeatherData(weatherData)
}
