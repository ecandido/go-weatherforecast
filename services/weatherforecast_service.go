package services

import (
	"go-weatherforecast/entities"
	"go-weatherforecast/repositories"
)

type WeatherForecastService interface {
	GetWeatherForecasts() []entities.WeatherForecast
}

type weatherForecastService struct {
	repository repositories.WeatherForecastRepository
}

func NewWeatherForecastService(repo repositories.WeatherForecastRepository) WeatherForecastService {
	return &weatherForecastService{repository: repo}
}

func (s *weatherForecastService) GetWeatherForecasts() []entities.WeatherForecast {
	return s.repository.GetAll()
}
