package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/rivo/tview"
)

const fileName = "todolist.csv"

func tui() {
	app := tview.NewApplication()

	table := tview.NewTable().SetBorders(true).SetSelectable(true, false)

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("error opening the file: ", err)
	}

	defer file.Close()
	// updates
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Errorf("error reading the file: %v", err)
	}

	loadTable(table, records)

	addTask := tview.NewInputField().
		SetLabel("New Task: ").
		SetFieldWidth(50)

	textView := tview.NewTextView()

	addTaskButton := tview.NewButton("Add").SetSelectedFunc(
		func() {
			task := addTask.GetText()
			if len(task) <= 3 {
				go func() {
					textView.SetText("Invalid Task")
					time.Sleep(3 * time.Second)
					app.QueueUpdateDraw(func() {
						textView.SetText("")
					})
				}()
			} else {
				err = createTask(task, fileName)
				if err != nil {
					fmt.Println(err)
				}
				records, err = readTasks(fileName)
				if err != nil {
					fmt.Println(err)
				}
				loadTable(table, records)

				go func() {
					textView.SetText("Task Successfully added")
					time.Sleep(3 * time.Second)
					app.QueueUpdateDraw(func() {
						textView.SetText("")
					})
				}()

			}
			// add a confirmation messge or refresh ui to add

		})

	rightColumn := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(addTask, 0, 1, false).
		AddItem(textView, 0, 2, false).
		AddItem(addTaskButton, 3, 0, false)

	flex := tview.NewFlex().
		AddItem(table, 80, 1, true).
		AddItem(rightColumn, 0, 2, false)

	flex.SetBorder(true).SetTitle("  ToDoList  ").SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
