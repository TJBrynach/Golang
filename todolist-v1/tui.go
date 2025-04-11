package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/rivo/tview"
)

const fileName = "todolist.csv"

func loadTable(table *tview.Table, records [][]string) {

	table.Clear()

	for rowIndex, row := range records {
		for colIndex, cell := range row {
			if colIndex != 0 {
				// need to set up logic to only show a row if its not completed
				tablecell := tview.NewTableCell(cell).
					SetAlign(tview.AlignCenter)

				table.SetCell(rowIndex, colIndex, tablecell)
			}

		}
	}

}

func tui() {
	app := tview.NewApplication()

	table := tview.NewTable().SetBorders(true).SetSelectable(true, true)

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("error opening the file: ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Errorf("error reading the file: %v", err)
	}

	loadTable(table, records)

	addTask := tview.NewInputField().
		SetLabel("New Task: ").
		SetFieldWidth(50)

	addTaskButton := tview.NewButton("Add").SetSelectedFunc(
		func() {
			task := addTask.GetText()
			err = createTask(task, fileName)
			if err != nil {
				fmt.Println(err)
			}
			records, err = readTasks(fileName)
			if err != nil {
				fmt.Println(err)
			}
			loadTable(table, records)
			// add a confirmation messge or refresh ui to add

		})

	rightColumn := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(addTask, 0, 2, false).
		AddItem(addTaskButton, 3, 0, false)

	flex := tview.NewFlex().
		AddItem(table, 80, 1, true).
		AddItem(rightColumn, 0, 2, false)

	flex.SetBorder(true).SetTitle("  ToDoList  ").SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
