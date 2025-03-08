package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		fmt.Print(string(r.Body))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Print("successfully")

	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("foundLink: %q -> %s\n", e.Text, link)
	})
	c.Visit("https://skills-todo-api.eliaschen.dev/Task")
}
