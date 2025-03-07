package main

import "time"

type Expense struct {
	ID     string    `json:"id"`
	Item   string    `json:"item"`
	Amount float32   `json:"amount"`
	Date   time.Time `json:"date"`
}

// func addExpense(category string, amount float32) (Expense, error) {

// }
