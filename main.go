package main

import (
	"go-weatherforecast/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/weatherforecast", controllers.GetWeatherForecast)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
