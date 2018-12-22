package main

import (
	"github.com/gocolly/colly"
)

func main() {

	scrapPage("https://www.moex.com/ru/contract.aspx?code=BR-1.19")

}

func scrapPage(url string) {

	c := colly.NewCollector()

	// find table and it's rows and columns
	c.OnHTML("table.contract-open-positions > tbody > tr", func(e *colly.HTMLElement) {

		content := e.DOM.Find("td:nth-child(2)").Text()

		println(content)
	})

	c.Visit(url)

}
