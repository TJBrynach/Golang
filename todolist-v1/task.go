package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        string
	Title     string
	Completed bool
	CreatedAt time.Time
}

// refactored csv into const
const (
	colID = iota
	colTitle
	colCompleted
	colCreatedAt
)

func (t Task) wrappedTitle() string {

	var finalList []string

	words := strings.Split(t.Title, " ")

	realIndex := 0

	for index, word := range words {
		m := index % 3

		if m == 2 {
			finalList = append(finalList, word)
			finalList = append(finalList, "\n")
		} else {
			finalList = append(finalList, word)
			realIndex++
		}
	}
	finalString := strings.Join(finalList, " ")
	return finalString
}

func (t Task) Display() string {
	status := "X"
	if t.Completed {
		status = "✅"
	}
	return fmt.Sprintf("[%s] %s (Created: %s)", status, t.Title, t.CreatedAt.Format("2023-15-10 15:04:23"))
}

func createTask(title string, fileName string) error {
	newTask := Task{
		ID:        uuid.New().String(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	tasks, err := readTasks(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	count := 0
	for _, task := range tasks {
		if newTask.Title == task.Title && task.Completed == false {
			count++
		}
	}
	if count == 0 {
		if err := saveToCSV(newTask, fileName); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("duplicate task")
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
		if i == colID || record[colTitle] != title {
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
		if i == colID || record[colTitle] != title {
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
