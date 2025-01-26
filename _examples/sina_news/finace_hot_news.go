package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("weibo.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	// Find and visit all links
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	e.Request.Visit(e.Attr("href"))
	//})

	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
		//fmt.Println("Content", r.Body)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Content", r.Body)
	})

	c.Visit("https://weibo.com/hot/search")
}
