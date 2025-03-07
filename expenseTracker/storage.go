package main

import (
	"encoding/json"
	"fmt"
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
	if err != nil && os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	defer file.Close()

	var expenses []Expense

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&expenses); err != nil {
		return nil, fmt.Errorf("erroring on decode: %v", err)
	}

	fmt.Println("data: ")
	fmt.Println(expenses)
	return expenses, nil
}

func writeJSON(expense Expense) error {
	fileName := "expenses.json"
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

// func writeJSON(expense Expense) error {
// 	fileName := "expenses.json"

// 	_, err := os.Stat(fileName)
// 	if err != nil && os.IsNotExist(err) {
// 		return fmt.Errorf("error checking file : %v", err)

// 	if err == nil {
// 		var expenses []Expense

// 		data, err := json.MarshalIndent(expenses, "", "  ")
// 		if err != nil {
// 			return fmt.Errorf("failed to Marshal blank data: %v", err)
// 		}

// 		err = os.WriteFile(fileName, data, 0644)
// 		if err != nil {
// 			return fmt.Errorf("failed to save file, %v", err)
// 		}
// 		fmt.Println("successfully created JSON")
// 	}

// 	}
// 	// if the file exists.

// 	data, err := json.MarshalIndent(expense, "", "  ")
// 	if err != nil {
// 		return fmt.Errorf("error Marshalling data")
// 	}

// 	err = os.WriteFile(fileName, data, 0644)
// 	if err != nil {
// 		return fmt.Errorf("error writing to file")
// 	}
// 	fmt.Println("successfully updated file")

// 	}
// 	return nil
// }

// func readJSON() {}
