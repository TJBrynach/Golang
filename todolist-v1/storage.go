package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
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

func listTasks(fileName string) ([]Task, error) {
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

	var tasks []Task

	for i, record := range records {
		if i == 0 {
			continue //Skip header
		}
		completed := record[2] == "true"
		createdAt, _ := time.Parse(time.RFC3339, record[3])

		task := Task{
			ID:        record[0],
			Title:     record[1],
			Completed: completed,
			CreatedAt: createdAt,
		}
		tasks = append(tasks, task)

	}
	return tasks, nil
}
