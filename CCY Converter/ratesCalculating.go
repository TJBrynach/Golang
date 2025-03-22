package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const API_URL = "https://v6.exchangerate-api.com/v6/c5492bb0169c65d93a2cdf88/latest/"

type ExchangeRates struct {
	Base            string             `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
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

	err = json.Unmarshal([]byte(body), &rates)
	if err != nil {
		return rates, err
	}

	return rates, nil
}
