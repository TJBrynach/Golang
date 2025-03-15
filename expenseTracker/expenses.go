package main

import (
	"fmt"
	"time"
)

type Expense struct {
	ID     string    `json:"id"`
	Item   string    `json:"item"`
	Amount float32   `json:"amount"`
	Date   time.Time `json:"date"`
}

func addExpense(item string, amount float32) error {

	new_expense := Expense{
		Item:   item,
		Amount: amount,
		Date:   time.Now(),
	}
	// fileName := "expenses.json"
	// content, err := readJSON(fileName)
	// if err != nil {
	// 	return []Expense{}, fmt.Errorf("errored reading file: %v", err)
	// }

	err := writeJSON(new_expense)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	// read expenses into go struct
	// create struct item from data
	// write expense into json
	return nil
}
