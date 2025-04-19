package main

import (
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/net/html"
)

func fetchTitle(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf(" Failed to fetch for %s: %v\n", url, err)
		return
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf(" Failed to parse HTML for %s: %v\n", url, err)
		return
	}

	var title string

	// the package parses a html document into a tree of nodes - expressed by *html.Node
	// each node has a type (html.ElementNode or html.TextNode)
	// Data - name of tag, title/div
	// FirstChild - first nested elemnt or text node
	// NextSibling - points ot next node on the same level

	// n.Type == html.ElementNode -> Is this node an actual html tag or just text?
	// && n.Data == "title" - an actual html tag and the name of the tag is title and it has contents ie text
	// finding <title>This is the page title</title>

	var findTitle func(*html.Node) // declare a recursive function that finds the title

	findTitle = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findTitle(c)
		}
	}

	findTitle(doc)

	fmt.Printf("%s - %s\n", url, title)

}

func main() {
	urls := []string{"https://www.bbc.co.uk/", "https://www.facebook.com/", "https://www.youtube.com/", "https://www.reddit.com/"}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchTitle(url, &wg)
	}

	wg.Wait()
	fmt.Println("All titles retrieved")
}
