package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type Task struct {
	ID        string
	Title     string
	Completed bool
	CreatedAt time.Time
}

func (t Task) Display() string {
	status := "X"
	if t.Completed {
		status = "✅"
	}
	return fmt.Sprintf("[%s] %s (Created: %s)", status, t.Title, t.CreatedAt.Format("2023-15-10 15:04:23"))
}

func createTask(title string, fileName string) error {
	task := Task{
		ID:        uuid.New().String(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	if err := saveToCSV(task, fileName); err != nil {
		return err
	}
	return nil
}

func deleteTasks(title string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}

	var updatedRecords [][]string
	taskDeleted := false

	for i, record := range records {
		if i == 0 || record[1] != title {
			updatedRecords = append(updatedRecords, record)
		} else {
			taskDeleted = true
		}
	}

	if !taskDeleted {
		return fmt.Errorf("task %s not found", err)
	}

	file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open CSV for writing: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.WriteAll(updatedRecords); err != nil {
		return fmt.Errorf("failed to write updated records to the file due to: %v", err)
	}

	fmt.Println("Task deleted successfully")
	return nil
}

func completeTask(title string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error accessing the file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	var updatedRecords [][]string

	for i, record := range records {
		if i == 0 || record[1] != title {
			updatedRecords = append(updatedRecords, record)
		} else {
			record[2] = "true"
			updatedRecords = append(updatedRecords, record)

		}
	}
	file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening the file to write: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Flush()

	if err := writer.WriteAll(updatedRecords); err != nil {
		return fmt.Errorf("failed to write updates to the csv: %v", err)
	}

	fmt.Println(title, " marked as complete")
	return nil
}

func markTaskAsDone(table *tview.Table, row int) {

}
