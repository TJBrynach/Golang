package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// we use colly to scrape and store prices for products on amazon
// takes an amazon product url
// scrapes products price
// checks price against target value
// if price has dropped send notification

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("scraping: ", r.StatusCode)
	})
	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.amazon.co.uk/")

}
