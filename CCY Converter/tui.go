package main

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

func gui() {

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
			_, from := baseCurrency.GetCurrentOption()
			_, to := farCurrency.GetCurrentOption()
			amount, err := strconv.ParseFloat(amountField.GetText(), 64)
			if err != nil {
				resultField.SetText("[red]Invalid amount![-]")
			}
			result := convertCCY(rates, from, to, amount)
			resultField.SetText(fmt.Sprintf("[green]Converted:%s, %s, %.2f %.2f[-]", from, to, amount, result))
		})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(baseCurrency, 1, 0, true).
		AddItem(farCurrency, 1, 0, true).
		AddItem(amountField, 1, 0, true).
		AddItem(resultField, 1, 0, true).
		AddItem(button, 1, 0, true)

	flex.SetBorder(true).SetTitle("CCY Converter")
	// // _, selectedBase := form.GetFormItem(0).(*tview.DropDown).GetCurrentOption()

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
