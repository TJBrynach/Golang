package main

import (
	"fmt"
)

func main() {
	base := "USD"
	rates, err := getExchangeRates(base)
	if err != nil {
		return
	}

	for key, value := range rates.ConversionRates {
		fmt.Println(key, value)
	}
	gui()
}
