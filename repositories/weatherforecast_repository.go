package repositories

import (
	"go-weatherforecast/entities"
	"time"
)

type WeatherForecastRepository interface {
	GetAll() []entities.WeatherForecast
}

type weatherForecastRepository struct{}

func NewWeatherForecastRepository() WeatherForecastRepository {
	return &weatherForecastRepository{}
}

func (r *weatherForecastRepository) GetAll() []entities.WeatherForecast {
	return []entities.WeatherForecast{
		{Date: time.Now().AddDate(0, 0, 1), TempC: 25, Summary: "Hot"},
		{Date: time.Now().AddDate(0, 0, 2), TempC: 22, Summary: "Warm"},
		{Date: time.Now().AddDate(0, 0, 3), TempC: 18, Summary: "Cool"},
	}
}
