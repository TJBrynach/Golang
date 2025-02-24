package main

import (
	"fmt"
	"os"
)

func handleCLI() {
	fmt.Println("CLI started:")
	if len(os.Args) != 4 {
		fmt.Println("Usage: <operation> <num1> <num2>")
		return
	}
	action := os.Args[1]
	numeral1 := os.Args[2]
	numeral2 := os.Args[3]

	var result interface{}
	var err error

	switch action {
	case "add":
		fmt.Println("starting addition cli")
		result, err = addition(numeral1, numeral2)
	case "subtract":
		fmt.Println("starting subtraction")
		result, err = subtraction(numeral1, numeral2)
	case "multiply":
		fmt.Println("starting multiply")
		result, err = multiply(numeral1, numeral2)
	case "divide":
		fmt.Println("starting division")
		result, err = divide(numeral1, numeral2)
	default:
		fmt.Println("used the incorrect function, options are: add, subtract, multiply, divide")
		return
	}

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}

}
