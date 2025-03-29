package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const api_key = "ce9340b365a450eec13c60f1e28607f0"

type WeatherData struct {
	City string `json:"name"`
	Main struct {
		Humidity int `json:"humidity"`
	} `json:"main"`
}

func main() {
	var data WeatherData
	city := "london"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, api_key)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("hello", data)
		return
	} else {
		fmt.Println(err)
	}
	fmt.Println(data)
}
