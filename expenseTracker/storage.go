package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// create file works
func createJSON() error {
	fileName := "expenses.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		var expense []Expense
		data, err := json.MarshalIndent(expense, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshalling, %v", err)
		}

		err = os.WriteFile(fileName, data, 0644)
		if err != nil {
			return fmt.Errorf("error creating file, %v", err)
		}
		fmt.Println("file created")
	}
	return nil
}

func readJSON(fileName string) ([]Expense, error) {
	file, err := os.Open(fileName)
	if err != nil {
		// if file doesn't exist
		if os.IsNotExist(err) {
			err = createJSON()
			if err != nil {
				return nil, fmt.Errorf("error creating JSON file: %v", err)
			}
			return []Expense{}, nil
		}
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	if len(content) == 0 {
		return []Expense{}, nil
	}

	var expenses []Expense

	if err := json.Unmarshal(content, &expenses); err != nil {
		return nil, fmt.Errorf("erroring on decode: %v", err)
	}

	// fmt.Println("data: ")
	// fmt.Println(expenses)
	return expenses, nil
}

func updateJSON(expense Expense) error {
	fileName := "expenses.json"

	// need a check to see if the file is empty
	data, err := readJSON(fileName)
	if err != nil {
		return fmt.Errorf("failed to read JSON: %v", err)
	}
	data = append(data, expense)

	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error creating JSON, %v", err)
	}

	err = os.WriteFile(fileName, updatedJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %v", err)
	}
	return nil
}
