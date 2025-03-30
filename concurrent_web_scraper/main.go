package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	var wg sync.WaitGroup

	cities := []string{"london", "toronto", "singapore", "tokyo", "istanbul", "washington"}

	for _, city := range cities {
		wg.Add(1)

		c := city

		go func(c string) {
			defer wg.Done()
			RetrieveWeather(c)
		}(c)

	}
	wg.Wait()
	fmt.Println("complete", time.Since(startTime))
}
