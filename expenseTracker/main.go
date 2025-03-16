package main

import (
	"fmt"
)

func main() {

	err := addExpense("Ice cream", 2.99)
	if err != nil {
		fmt.Printf("error adding expense, %v", err)
	}

	// err = listExpenses()
	// if err != nil {
	// 	fmt.Println("error printing list")
	// }

	// err = calcExpenses()
	// if err != nil {
	// 	fmt.Println("error printing calculating expenses")
	// }

	// err = calcItemExpenses("bicycle")
	// if err != nil {
	// 	fmt.Println("error calculating items exepenses")
	// }

	count, err := countItemExpenses("ice cream")
	if err != nil {
		fmt.Println("error getting the count")
	}
	fmt.Println(count)
}
