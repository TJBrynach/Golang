package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp

}

func main() {
	a := 5
	b := 10
	fmt.Printf("a is %v, b is %v", a, b)
	swap(&a, &b)
	fmt.Println(" ")
	fmt.Printf("a is %v, b is %v", a, b)

}
