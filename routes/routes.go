package routes

import (
	"assignment-ketiga/controllers"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/weathers", controllers.GetWeather)

	http.ListenAndServe(":8080", nil)
}