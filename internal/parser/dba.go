package parser

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Dba struct {
}

func NewDba() Dba {
	return Dba{}
}

func (d *Dba) IsActive(url string) bool {
	fmt.Println("isActive called from parser.dba")
	var isActive bool

	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("span", func(_ int, el *colly.HTMLElement) {
			if strings.Contains(el.Text, "Skriv til sælger") {
				isActive = true
			}
		})
		e.ForEach("div", func(_ int, el *colly.HTMLElement) {
			if strings.Contains(el.Text, "Denne annonce er inaktiv") {
				isActive = false
			}
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	return isActive
}
