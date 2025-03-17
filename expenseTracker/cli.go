package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HandleCLI() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nExpense Tracker menu:")
		fmt.Println("1. Add expense")
		fmt.Println("2. View expenses")
		fmt.Println("3. Calcuate expenses")
		fmt.Println("4. Total spent on specific item")
		fmt.Println("5. Delete expense")
		fmt.Println("6. Exit")

		fmt.Println("Enter choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("We can add an expense, can you provide more details; name and cost")
			err := addExpense(reader)
			if err != nil {
				return
			}
			fmt.Println("Successfully added item")
		case "2":
			fmt.Println(" viewing expenses")
			err := listExpenses()
			if err != nil {
				return
			}
		case "3":
			fmt.Println(" calcing expenses")
			err := calcExpenses()
			if err != nil {
				return
			}

		case "4":
			fmt.Println(" calcing items expense")
			err := calcItemExpenses(reader)
			if err != nil {
				return
			}
		case "5":
			fmt.Println("deleting expense")
		case "6":
			fmt.Println(" exiting")
			return
		default:
			fmt.Println(" invalid option")
		}

		fmt.Println("would you like to return to the main menu? (y/n)")
		newInput, _ := reader.ReadString('\n')
		newInput = strings.TrimSpace(newInput)
		if newInput == "y" {
			continue
		} else {
			return
		}
	}
}
