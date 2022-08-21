package parser

import (
	"Projects/WordAnalytics/logger"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type (
	urlS = string
)

var text = ""

func Parse(url urlS) string {
	c := colly.NewCollector()
	logging := logger.GetLogger()

	logging.Info("Find and visit all links")
	c.OnHTML("a", onHtmlCallback)
	c.OnHTML("h1", onHtmlCallback)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		text = ""
	})

	c.Visit(url)

	text = strings.ToLower(text)

	return text
}

func onHtmlCallback(e *colly.HTMLElement) {
	text += " " + strings.TrimSpace(e.Text)
}
