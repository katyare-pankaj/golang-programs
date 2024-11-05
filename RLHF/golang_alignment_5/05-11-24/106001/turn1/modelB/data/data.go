package data

type Weather struct {
	City        string `json:"city"`
	Temperature int    `json:"temperature"`
}

func GetWeather() *Weather {
	return &Weather{
		City:        "New York",
		Temperature: 25,
	}
}
