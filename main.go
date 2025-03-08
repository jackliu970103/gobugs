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

}
