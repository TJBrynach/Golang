package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// we use colly to scrape and scrape tmdb

type Film struct {
	Title       string
	UserScore   int
	Director    string
	Description string
	ReleaseYear int
}

func main() {
	var (
		movies []Film
		// mu     sync.Mutex
	)
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

	c.OnHTML("a.image", func(r *colly.HTMLElement) {
		r.Request.Visit(r.Attr("href"))

	})

	c.OnHTML("span.release_date", func(e *colly.HTMLElement) {
		year := e.Text
		fmt.Println(year)

	})

	c.Visit("https://www.themoviedb.org/movie")

	for _, movie := range movies {
		fmt.Println(movie.Title, movie.Director)
	}

}
