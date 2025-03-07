package main

import (
	"fmt"
	"time"
)

type Expense struct {
	ID          string    `json:"id"`
	Category    string    `json:"category"`
	Amount      float32   `json:"amount"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// func createExpense(category string, amount float32, description string) error {

// }

func main() {
	// expense := Expense{
	// 	Category:    "food",
	// 	Amount:      12.23,
	// 	Description: "bananas and baked beans",
	// }
	// err := writeJSON(expense)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	err := createJSON()
	if err != nil {
		fmt.Println(err)
	}
}
