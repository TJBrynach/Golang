package main

import (
	"fmt"
	"strings"
	"time"
)

type Expense struct {
	ID     string    `json:"id"`
	Item   string    `json:"item"`
	Amount float32   `json:"amount"`
	Date   time.Time `json:"date"`
}

// better way to display each expense
func (e Expense) Display() string {
	return fmt.Sprintf("%s: £%.2f", e.Item, e.Amount)
}

func (e Expense) IdIncrement() int {}

func addExpense(item string, amount float32) error {

	new_expense := Expense{
		Item:   item,
		Amount: amount,
		Date:   time.Now(),
	}
	err := updateJSON(new_expense)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}

func listExpenses() error {
	content, err := readJSON("expenses.json")
	if err != nil {
		fmt.Printf("error reading file to list: %v", err)
	}
	fmt.Println("Here are your expenses:")
	for _, expense := range content {
		fmt.Println(expense.Display())
	}

	return nil
}

func calcExpenses() error {
	content, err := readJSON("expenses.json")
	if err != nil {
		return fmt.Errorf("failed to read json to calc: %v", err)
	}
	var totalExpenses float32
	for _, e := range content {
		fmt.Println(e.Amount)
		totalExpenses += e.Amount
	}
	fmt.Printf("Total Expenses: %.2f\n", totalExpenses)
	return nil
}

func calcItemExpenses(item string) error {
	const fileName = "expenses.json"
	content, err := readJSON(fileName)
	if err != nil {
		return fmt.Errorf("return x")
	}
	// set to lower case
	lowerItem := strings.ToLower(item)

	var itemTotalExpenses float32
	for _, e := range content {
		if strings.ToLower(e.Item) == lowerItem {
			itemTotalExpenses += e.Amount
		}
	}

	fmt.Printf("Total expenses for %v: %.2f", item, itemTotalExpenses)
	return nil
}

func countItemExpenses(item string) (int16, error) {
	content, err := readJSON("expenses.json")
	if err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	// to lower
	lowerItem := strings.ToLower(item)
	var counter int16
	for _, e := range content {
		if strings.ToLower(e.Item) == lowerItem {
			counter += 1
		}
	}
	return counter, nil
}

func deleteExpense(item string) error {
	//
	content, err := readJSON("expenses.json")
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	lowerItem := strings.ToLower(item)

	for _, e := range content {
	}

	return nil
}
