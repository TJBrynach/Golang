package main

import "github.com/rivo/tview"

func gui() {
	// box := tview.NewBox().SetBorder(true).SetTitle(" CCY Converter ")

	rates, err := getExchangeRates()
	if err != nil {
		return
	}

	app := tview.NewApplication()

	baseCurrency := tview.NewDropDown().
		SetLabel("From Currency: ").
		SetOptions(ccyOptions(rates), nil)

	farCurrency := tview.NewDropDown().
		SetLabel("To Currency: ").
		SetOptions(ccyOptions(rates), nil)

	amountField := tview.NewInputField().
		SetLabel("Amount: ").
		SetFieldWidth(10)

	resultField := tview.NewTextView().
		SetLabel("Converted Amount: ").
		SetDynamicColors(true)

	button := tview.NewButton("Convert").
		SetSelectedFunc(func() {
			from, _ := baseCurrency.GetCurrentOption()
			to, _ := farCurrency.GetCurrentOption()
		})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(baseCurrency, 1, 0, true).
		AddItem(farCurrency, 1, 0, true).
		AddItem(amountField, 1, 0, true).
		AddItem(resultField, 1, 0, true).
		AddItem(button, 1, 0, true)

	// // bases := []string{"GBP", "USD", "SGD", "EUR", "NOK", "SEK", "HKD"}
	// // // var selectedBase string

	// // // form := tview.NewForm().
	// // // 	AddDropDown("Select Base currency (hit enter): ", bases, 0, nil).
	// // // 	AddInputField("Amount in base currency to be converted: ", "", 20, nil, nil).
	// // // 	AddButton("Convert", getRate(rates, "EUR")).
	// // // 	AddButton("Quit", func() { app.Stop() })
	// // // form.SetBorder(true).SetTitle("CCY Converter").SetTitleAlign(tview.AlignLeft)

	// // _, selectedBase := form.GetFormItem(0).(*tview.DropDown).GetCurrentOption()

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	// base := "USD"

	// far := "GBP"

	// for key, value := range rates.ConversionRates {
	// 	if key == far {
	// 		fmt.Println(selectedBase, key, value)
	// 	}

	// }
}
