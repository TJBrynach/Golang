package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WeatherData struct {
	City string `json:"name"`
	Main struct {
		Humidity int `json:"humidity"`
	} `json:"main"`
}

const api_key = "ce9340b365a450eec13c60f1e28607f0"

func RetrieveWeather(city string) (*WeatherData, error) {
	var data WeatherData
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, api_key)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching weather for:%s, %v\n", city, err)
		return &data, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Non-OK HTTP status for %s: %s\n", city, resp.Status)
		return &data, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding weather data for:%s, %v\n", city, err)
		return &data, err
	}
	// fmt.Printf("The humidity if %s is %v \n", data.City, data.Main.Humidity)
	return &data, nil
}
