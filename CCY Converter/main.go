package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const API_URL = "https://v6.exchangerate-api.com/v6/c5492bb0169c65d93a2cdf88/latest/"

type ExchangeRates struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

func getExchangeRates(base string) (ExchangeRates, error) {
	var rates ExchangeRates
	fmt.Println(API_URL + base)
	resp, err := http.Get(API_URL + base)
	if err != nil {
		return rates, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return rates, err
	}

	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(body), &rates)
	if err != nil {
		return rates, err
	}
	fmt.Println(jsonMap)
	fmt.Println(rates)
	// fmt.Println(jsonMap["conversion_rate"])
	// fmt.Println(jsonMap["target_code"])
	// fmt.Println(jsonMap["base_code"])

	return rates, nil
}

func main() {
	// const key = "c5492bb0169c65d93a2cdf88"
	// client := resty.New()

	// resp, err := client.R().Get("https://v6.exchangerate-api.com/v6/c5492bb0169c65d93a2cdf88/pair/EUR/GBP")
	base := "USD"
	rates, err := getExchangeRates(base)
	if err != nil {
		return
	}

	fmt.Println("Exchange rates: ", rates.Rates)
}

// {"result":"success","documentation":"https://www.exchangerate-api.com/docs","terms_of_use":"https://www.exchangerate-api.com/terms","time_last_update_unix":1742169601,"time_last_update_utc":"Mon, 17 Mar 2025 00:00:01 +0000","time_next_update_unix":1742256001,"time_next_update_utc":"Tue, 18 Mar 2025 00:00:01 +0000","base_code":"EUR","target_code":"GBP","conversion_rate":0.8413}
