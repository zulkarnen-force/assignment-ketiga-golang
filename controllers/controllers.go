package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Result map[string]interface{} `json:"result"`
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	bytesFile, err := ioutil.ReadFile("weather.json")
	
	if err != nil {
		fmt.Println(err.Error())
	}

	var weatherData map[string]interface{}

	json.Unmarshal(bytesFile, &weatherData)
	
	res := Response {
		Result: weatherData,
	}

	json.Marshal(&res)

	    //Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(res)
}