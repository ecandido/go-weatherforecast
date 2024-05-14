package entities

import (
	"time"
)

type WeatherForecast struct {
	Date    time.Time `json:"date"`
	TempC   int       `json:"temperatureC"`
	Summary string    `json:"summary"`
}
