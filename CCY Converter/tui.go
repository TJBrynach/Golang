package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func gui() {
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}

	base := "USD"
	rates, err := getExchangeRates(base)
	if err != nil {
		return
	}

	for key, value := range rates.ConversionRates {
		fmt.Println(key, value)
	}
}
