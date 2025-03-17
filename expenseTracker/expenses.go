package main

import (
	"bufio"
	"fmt"
	"strconv"
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

func IdIncrement() string {
	content, err := readJSON("expenses.json")
	if err != nil {
		return "0"
	}

	var id_num int
	if len(content) > 0 {
		maxID, err := strconv.Atoi(content[len(content)-1].ID)
		if err != nil {
			fmt.Println("erroring getting maxid")
			return "0"
		}
		id_num = maxID + 1
	} else {
		id_num = 1
	}
	return strconv.Itoa(id_num)
}

func addExpense(reader *bufio.Reader) error {

	new_id := IdIncrement()

	fmt.Println("Enter item description: ")
	description, _ := reader.ReadString('\n')
	item := strings.TrimSpace(description)

	fmt.Println("How much did it cost: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)

	amount64, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return fmt.Errorf("invalid amount. Please enter a valid number.")
	}
	amount := float32(amount64)

	new_expense := Expense{
		ID:     new_id,
		Item:   item,
		Amount: amount,
		Date:   time.Now(),
	}

	err = updateJSON(new_expense)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Println(new_expense.Display())
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
		totalExpenses += e.Amount
	}
	fmt.Printf("Total Expenses: £%.2f\n", totalExpenses)
	return nil
}

func calcItemExpenses(reader *bufio.Reader) error {
	const fileName = "expenses.json"
	content, err := readJSON(fileName)
	if err != nil {
		return fmt.Errorf("return x")
	}
	fmt.Println("Which item would you like calculate expenses for?")
	item, _ := reader.ReadString('\n')
	item = strings.TrimSpace(item)
	// set to lower case
	lowerItem := strings.ToLower(item)

	var itemTotalExpenses float32
	for _, e := range content {
		if strings.ToLower(e.Item) == lowerItem {
			itemTotalExpenses += e.Amount
		}
	}

	fmt.Printf("Total expenses for %v: £%.2f", item, itemTotalExpenses)
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

func summary() error {
	content, err := readJSON("expenses.json")
	if err != nil {
		return fmt.Errorf("error reading json to summarise:%v", err)
	}
	m := make(map[string]float32)
	for _, e := range content {
		_, exists := m[e.Item]
		if exists == false {
			m[e.Item] = float32(e.Amount)
		} else {
			m[e.Item] += float32(e.Amount)
		}
	}
	for item, amount := range m {
		fmt.Println(item, ": £", amount)
	}
	// loop through and count per item
	return nil
}

func deleteExpense(reader *bufio.Reader) error {
	//
	content, err := readJSON("expenses.json")
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// lowerItem := strings.ToLower(item)

	// for _, e := range content {
	// }

	return nil
}
