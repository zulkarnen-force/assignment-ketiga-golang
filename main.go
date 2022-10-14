package main

import (
	"assignment-ketiga/models"
	"assignment-ketiga/routes"
)


func main() {
	weather := models.Weather{}
	go weather.RealTimeWeather()
	routes.StartServer()
}