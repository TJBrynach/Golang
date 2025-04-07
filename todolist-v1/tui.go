package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/rivo/tview"
)

func tui() {
	app := tview.NewApplication()

	table := tview.NewTable().SetBorders(true).SetSelectable(true, true)

	file, err := os.Open("todolist.csv")

	if err != nil {
		fmt.Println("error opening the file: ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Errorf("error reading the file: %v", err)
	}
	for rowIndex, row := range records {
		for colIndex, cell := range row {
			tablecell := tview.NewTableCell(cell).
				SetAlign(tview.AlignCenter)

			table.SetCell(rowIndex, colIndex, tablecell)
		}
	}

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 1, true)

	flex.SetBorder(true).SetTitle("  ToDoList  ").SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
