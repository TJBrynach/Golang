package main

import "github.com/gocolly/colly"

// we use colly to scrape and store prices for products on amazon
// takes an amazon product url
// scrapes products price
// checks price against target value
// if price has dropped send notification

func main() {
	c := colly.NewCollector()

}
https://go-colly.org/docs/examples/basic/