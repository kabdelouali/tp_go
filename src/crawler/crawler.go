//a concurrent web crawler who explore the link graph of the web in breadth-first order
//To test, go run crawler.go https://www.catie.fr
package main

import (
	"fmt"
	"net/http"
	"os"

	html "golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	resp, errr := http.Get(url)
	if errr != nil {
		fmt.Println("ERROR")
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("ERROR")
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					links = append(links, a.Val)
					break
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//recursive call
		forEachNode(c, pre, nil)
	}
	if post != nil {
		post(n)
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		fmt.Println("ERROR")
	}
	return list
}

//!+
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	seen := make(map[string]bool)
	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()
	// unseenLinks <- (<-worklist)[0]

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				links := crawl(link)
				go func() { worklist <- links }()

			}
		}()
	}

	// // The main goroutine de-duplicates worklist items
	// // and sends the unseen ones to the crawlers.
	for links := range worklist {
		for _, link := range links {
			if seen[link] != true {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
