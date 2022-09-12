package parser

import (
	"Projects/WordAnalytics/internal/counter"
	"Projects/WordAnalytics/pkg/logger"
	"encoding/json"
	"github.com/gocolly/colly"
	"net/url"
	"strings"
)

var text = ""

func Parse(url string) string {
	c := colly.NewCollector()
	log := logger.GetLogger()

	log.Info("Find and visit all links")
	c.OnHTML("a", onHtmlCallback)
	c.OnHTML("h1", onHtmlCallback)

	c.OnRequest(func(r *colly.Request) {
		log.Info("Visiting", r.URL)
		text = ""
	})

	c.Visit(url)

	text = strings.ToLower(text)

	return text
}

func onHtmlCallback(e *colly.HTMLElement) {
	text += " " + strings.TrimSpace(e.Text)
}

func IsUrl(site string) bool {
	u, err := url.Parse(site)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func FindResult(jsonWords []byte, wordFromTg string) int {
	log := logger.GetLogger()
	var words []counter.Word

	err := json.Unmarshal(jsonWords, &words)
	if err != nil {
		log.Errorf("Unmarshal error")
	}

	for _, el := range words {
		if wordFromTg == el.Word {
			return el.Count
		} else {
			continue
		}
	}
	return 0
}
