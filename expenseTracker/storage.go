package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func createJSON() error {
	fileName := "expenses.json"
	_, err := os.Stat(fileName)
}

func writeJSON(expense Expense) error {
	fileName := "expenses.json"

	_, err := os.Stat(fileName)
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("error checking file : %v", err)

	if err == nil {
		var expenses []Expense

		data, err := json.MarshalIndent(expenses, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to Marshal blank data: %v", err)
		}

		err = os.WriteFile(fileName, data, 0644)
		if err != nil {
			return fmt.Errorf("failed to save file, %v", err)
		}
		fmt.Println("successfully created JSON")
	}

	} 
	// if the file exists.
	

	data, err := json.MarshalIndent(expense, "", "  ")
	if err != nil {
		return fmt.Errorf("error Marshalling data")
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file")
	}
	fmt.Println("successfully updated file")

	}
	return nil
}

// func readJSON() {}
