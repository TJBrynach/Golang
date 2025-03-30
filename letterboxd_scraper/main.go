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
	// c.Visit("https://www.amazon.co.uk/Amazon-streaming-device-supports-Vision/dp/B0CJL4J6FG?ref=dlx_20564_dg_dcl_B0CJL4J6FG_dt_mese3_cc_pi&pf_rd_r=E4H2V4RM05DGTYZH7YHV&pf_rd_p=b5b598aa-de73-4068-82b4-c1b94b3594cc")

	fmt.Println("visited site - now what")

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("scraping: ", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("❌ Request failed:", err)
	})

	c.OnHTML("html", func(r *colly.HTMLElement) {
		fmt.Println("there is text")
		// whole_value := r.FirstChild.Data
		// fmt.Println(whole_value)
	})
	c.Visit("https://letterboxd.com/")
}
