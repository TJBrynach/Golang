package main

import (
	"errors"
	"fmt"
	"strconv"
)

func toInt(a string) (int, error) {
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, errors.New("invalid number: " + a)
	}
	return n, nil
}

func addition(a string, b string) (int, error) {
	fmt.Println("starting addition function")
	numeralA, errA := toInt(a)
	numeralB, errB := toInt(b)
	if errA != nil || errB != nil {
		return 0, errors.New("invalid input")
	}
	return numeralA + numeralB, nil
}

func subtraction(a string, b string) (int, error) {
	fmt.Println("starting subtraction function")
	numeralA, errA := toInt(a)
	numeralB, errB := toInt(b)
	if errA != nil || errB != nil {
		return 0, errors.New("invalid input")
	}
	sum := numeralA - numeralB
	return sum, nil
}

func multiply(a string, b string) (int, error) {
	fmt.Println("starting multiplication function")
	numeralA, errA := toInt(a)
	numeralB, errB := toInt(b)
	if errA != nil || errB != nil {
		return 0, errors.New("invalid input")
	}
	sum := numeralA * numeralB
	return sum, nil

}

func divide(a string, b string) (float64, error) {
	fmt.Println("starting division function")
	numeralA, errA := strconv.ParseFloat(a, 64)
	numeralB, errB := strconv.ParseFloat(b, 64)
	if errA != nil || errB != nil {
		return 0.0, errors.New("invalid input")
	}
	if numeralB != 0 {
		sum := numeralA / numeralB
		return sum, nil
	} else {
		return 0.0, errors.New("can't divide by zero")
	}
}
