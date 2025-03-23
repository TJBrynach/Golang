package main

import (
	"fmt"
)

func worker(id int, ch chan string) {
	ch <- fmt.Sprintf("Worker %d done", id)
}

func main() {
	// initiate unbuffered int channel
	ch := make(chan string)

	for i := 0; i <= 3; i++ {
		go worker(i, ch)
	}

	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch)
	}
}
