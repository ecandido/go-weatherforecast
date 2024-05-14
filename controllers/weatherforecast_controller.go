package controllers

import (
	"encoding/json"
	"go-weatherforecast/repositories"
	"go-weatherforecast/services"
	"net/http"
)

func GetWeatherForecast(w http.ResponseWriter, r *http.Request) {

	repository := repositories.NewWeatherForecastRepository()
	service := services.NewWeatherForecastService(repository)
	forecasts := service.GetWeatherForecasts()

	json.NewEncoder(w).Encode(forecasts)
}
