package main

import (
	"fmt"
	"log"
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
			weatherData, err := RetrieveWeather(c)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("The humidity if %s is %v \n", weatherData.City, weatherData.Main.Humidity)
		}(c)

	}
	wg.Wait()
	fmt.Println("complete", time.Since(startTime))
}
