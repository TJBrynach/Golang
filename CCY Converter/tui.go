package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func gui() {
	// box := tview.NewBox().SetBorder(true).SetTitle(" CCY Converter ")
	app := tview.NewApplication()

	bases := []string{"GBP", "USD", "SGD", "EUR", "NOK", "SEK", "HKD"}
	// var selectedBase string

	form := tview.NewForm().
		AddDropDown("Select Base currency (hit enter): ", bases, 0, nil).
		AddInputField("Amount in base currency to be converted: ", "", 20, nil, nil).
		AddButton("Convert", nil).
		AddButton("Quit", func() { app.Stop() })
	form.SetBorder(true).SetTitle("CCY Converter").SetTitleAlign(tview.AlignLeft)

	_, selectedBase := form.GetFormItem(0).(*tview.DropDown).GetCurrentOption()

	if err := app.SetRoot(form, true).SetFocus(form).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	// base := "USD"
	rates, err := getExchangeRates(selectedBase)
	if err != nil {
		return
	}

	far := "GBP"

	for key, value := range rates.ConversionRates {
		if key == far {
			fmt.Println(selectedBase, key, value)
		}

	}
}
