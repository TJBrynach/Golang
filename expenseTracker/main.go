package main

import (
	"fmt"
	"time"
)

func main() {

	err := createJSON()
	if err != nil {
		fmt.Println(err)
	}

	expense := Expense{
		Item:   "Baked beans",
		Amount: 12.23,
		Date:   time.Now(),
	}

	err = writeJSON(expense)
	if err != nil {
		fmt.Println(err)
	}
}
