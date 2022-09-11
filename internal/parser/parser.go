package parser

import (
	"Projects/WordAnalytics/internal/counter"
	"Projects/WordAnalytics/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"strings"
)

var text = ""

func Parse(url string) string {
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

func IsUrl(site string) bool {
	u, err := url.Parse(site)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func FindResult(str []byte, word_from_tg string) int {
	log := logger.GetLogger()
	var words []counter.Word

	err := json.Unmarshal(str, &words)
	if err != nil {
		log.Errorf("Unmarshal error")
	}

	for _, el := range words {
		if word_from_tg == el.Word {
			return el.Count
		} else {
			continue
		}
	}
	return 0
}
