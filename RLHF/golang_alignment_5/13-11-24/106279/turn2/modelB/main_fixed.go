package main

const CELSIUS_TO_FAHRENHEIT_CONSTANT = 9.0 / 5.0

func celsiusToFahrenheitReadable(celsius float64) float64 {
	return (celsius * CELSIUS_TO_FAHRENHEIT_CONSTANT) + 32.0
}

func celsiusToFahrenheitEfficient(celsius float32) float32 {
	const (
		f1 = 9.0 / 5.0
		f2 = 32.0
	)

	return (celsius * float32(f1)) + float32(f2)
}
