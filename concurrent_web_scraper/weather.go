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

func RetrieveWeather(city string) {
	var data WeatherData
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, api_key)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("hello", data)
		return
	}
	fmt.Println(data)
}
