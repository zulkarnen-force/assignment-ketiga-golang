package models

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}


func (w Weather) writeJSON() {
	water := rand.Intn(10)
	wind := rand.Intn(10)

	weather := Weather{
		Water: water,
		Wind: wind,
	}

	file, _ := json.MarshalIndent(weather, "", " ")

	_ = ioutil.WriteFile("weather.json", file, 0644)

}

func (w Weather) RealTimeWeather() {

    for range time.Tick(time.Second * 15) {
		w.writeJSON()
    }

}