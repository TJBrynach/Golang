package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const fileName = "todolist.csv"

func tui() {
	app := tview.NewApplication()

	table := tview.NewTable().SetBorders(true).SetSelectable(true, false)

	//initialise textView
	textView := tview.NewTextView()
	textView.SetDynamicColors(true)

	file, err := os.Open(fileName)

	if err != nil {
		textView.SetText(fmt.Sprintf("[red]error opening the file: %v[-]", err))
	}

	defer file.Close()
	// updates
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		textView.SetText(fmt.Sprintf("[red]error reading the file: %v", err))
	}

	loadTable(table, records)

	addTask := tview.NewInputField().
		SetLabel("New Task: ").
		SetFieldWidth(50)

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'c':
				row, _ := table.GetSelection()
				task := table.GetCell(row, 0).Text
				completeTask(task, fileName)

				records, err = readTasks(fileName)
				if err != nil {
					fmt.Println(err)
				}
				loadTable(table, records)
				textView.SetText(fmt.Sprintf("[green]%v has been marked complete[-]", task))

				app.SetFocus(table)

				return nil
			case 'd':
				row, _ := table.GetSelection()
				task := table.GetCell(row, 0).Text
				deleteTasks(task, fileName)

				records, err = readTasks(fileName)
				if err != nil {
					fmt.Println(err)
				}

				loadTable(table, records)

				textView.SetText(fmt.Sprintf("[green]%v has been deleted[-]", task))

				app.SetFocus(table)

				return nil
			}
		}
		return event
	})

	addTaskButton := tview.NewButton("Add").SetSelectedFunc(
		func() {
			task := addTask.GetText()
			if len(task) <= 3 {
				go func() {
					textView.SetText("[red]Invalid Task[-]")
					time.Sleep(3 * time.Second)
					app.QueueUpdateDraw(func() {
						textView.SetText("")
					})
				}()
			} else {
				err = createTask(task, fileName)
				if err != nil {
					textView.SetText(fmt.Sprintf("[red]%v[-]", err))
				}
				records, err = readTasks(fileName)
				if err != nil {
					fmt.Println(err)
				}
				loadTable(table, records)

				addTask.SetText("")

				app.SetFocus(addTask)

				go func() {
					textView.SetText("[green]Task Successfully added[-]")
					time.Sleep(3 * time.Second)
					app.QueueUpdateDraw(func() {
						textView.SetText("")
					})
				}()

			}

		})

	rightColumn := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(addTask, 0, 1, false).
		AddItem(textView, 0, 2, false).
		AddItem(addTaskButton, 2, 0, false)

	flex := tview.NewFlex().
		AddItem(table, 0, 1, true).
		AddItem(rightColumn, 0, 1, false)

	flex.SetBorder(true).SetTitle("  ToDoList  ").SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
