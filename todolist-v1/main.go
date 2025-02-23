package main

import (
	"fmt"
)

func main() {
	const fileName = "todolist.csv"
	initCSV(fileName)

	if err := initCSV(fileName); err != nil {
		fmt.Println("Error initialising the csv file", err)
		return
	}
	HandleCLI()

}
