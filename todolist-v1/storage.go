package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rivo/tview"
)

func initCSV(fileName string) error {
	_, err := os.Stat(fileName) // check if file exists or not, _ is mute as we only care if it exists or not
	if os.IsNotExist(err) {     // checks the error for the it doesnt exist one - if it doesnt, we create a new file.
		file, err := os.Create(fileName) // creates a file
		if err != nil {
			return fmt.Errorf("failed to create CSV File: %v", err) // if error print error
		}
		defer file.Close() // prevents memory leaks & corruption

		// now we write to the file
		writer := csv.NewWriter(file) // creating the csv writer to allow writing
		defer writer.Flush()          // ensures all buffered data is written to the file.the writer buffers data before writing, so flushing ensures its saved

		headers := []string{"ID", "Title", "Completed", "CreatedAt"}
		if err := writer.Write(headers); err != nil {
			return fmt.Errorf("failed to write headers: %v", err)
		}
	}
	return nil // return on success
}

func saveToCSV(task Task, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close() // prevents memory leak

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// now add the record through the writer
	record := []string{
		task.ID,
		task.Title,
		fmt.Sprintf("%t", task.Completed),
		task.CreatedAt.Format(time.RFC3339),
	}

	if err := writer.Write(record); err != nil {
		return fmt.Errorf("failed to write to csv: %v", err)
	}
	return nil
}

func readTasks(fileName string) ([]Task, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	var tasks []Task
	for i, row := range records {
		// ignore header
		if i == 0 {
			continue
		}
		completed, _ := strconv.ParseBool(row[colCompleted])
		createdAt, _ := time.Parse(time.RFC3339, row[colCreatedAt])
		task := Task{
			ID:        row[colID],
			Title:     row[colTitle],
			Completed: completed,
			CreatedAt: createdAt,
		}

		tasks = append(tasks, task)
	}

	return tasks, nil

}

// func listTasks(fileName string) ([]Task, error) {
// 	records, err := readTasks(fileName)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var tasks []Task

// 	for i, record := range records {
// 		if i == 0 {
// 			continue //Skip header
// 		}
// 		completed := record[2] == "true"
// 		createdAt, _ := time.Parse(time.RFC3339, record[3])

// 		task := Task{
// 			ID:        record[colID],
// 			Title:     record[colTitle],
// 			Completed: completed,
// 			CreatedAt: createdAt,
// 		}
// 		tasks = append(tasks, task)

// 	}
// 	return tasks, nil
// }

// func loadTable(table *tview.Table, records [][]string) {

// 	table.Clear()

// 	realrowIndex := 0

// 	for _, row := range records {
// 		if row[colCompleted] == "true" {
// 			continue
// 		}

// 		visibleColIndex := 0
// 		for colIndex, cell := range row {
// 			if colIndex != colID && colIndex != colCompleted {
// 				// format date time
// 				if colIndex == colCreatedAt {
// 					var formattedCell string
// 					parsedTime, err := time.Parse(time.RFC3339, cell)
// 					if err == nil {
// 						formattedCell = parsedTime.Format("2006-01-02 15:04")
// 					} else {
// 						formattedCell = cell
// 					}
// 					tablecell := tview.NewTableCell(formattedCell).SetAlign(tview.AlignCenter)
// 					table.SetCell(realrowIndex, visibleColIndex, tablecell)
// 					visibleColIndex++
// 					//format title
// 				} else {
// 					tablecell := tview.NewTableCell(cell).
// 						SetAlign(tview.AlignCenter)
// 					table.SetCell(realrowIndex, visibleColIndex, tablecell)
// 					visibleColIndex++
// 				}
// 			}

// 		}
// 		realrowIndex++

// 	}
// }

func loadTasks(table *tview.Table, tasks []Task) {

	table.Clear()

	table.SetCell(0, 0, tview.NewTableCell("Title").SetAlign(tview.AlignCenter).SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell("CreatedAt").SetAlign(tview.AlignCenter).SetSelectable(false))

	realRowIndex := 1
	for _, task := range tasks {
		if task.Completed {
			continue
		}

		titleCell := tview.NewTableCell(task.wrappedTitle()).SetAlign(tview.AlignCenter)
		table.SetCell(realRowIndex, 0, titleCell)

		table.SetCell(realRowIndex, 1, tview.NewTableCell(task.CreatedAt.Format("2006-01-02 15:04")))

		realRowIndex++

	}
}
