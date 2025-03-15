package main

import (
	"fmt"
)

func main() {

	// err := createJSON()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// expense := Expense{
	// 	Item:   "Baked beans",
	// 	Amount: 12.23,
	// 	Date:   time.Now(),
	// }
	err := addExpense("Baked Beans", 2.25)
	// err := writeJSON(expense)
	if err != nil {
		fmt.Println("this is the error, %v", err)
	}
	// err := writeJSON()

}
