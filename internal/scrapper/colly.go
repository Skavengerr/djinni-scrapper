package scrapper

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func GetVacanciesByUrl(urls map[string]string) []Vacancy {
	var vacanciesSlice []Vacancy

	c := colly.NewCollector()

	handleHTML := func(e *colly.HTMLElement) {
		pattern := regexp.MustCompile(`\s+`)
		vacancyTitle := pattern.ReplaceAllString(e.Text, " ")
		vacancyURL := strings.TrimSpace(e.Request.AbsoluteURL(e.Attr("href")))

		vacancy := Vacancy{
			Title: strings.Trim(vacancyTitle, " "),
			URL:   vacancyURL,
		}
		vacanciesSlice = append(vacanciesSlice, vacancy)
	}

	for domain, url := range urls {
		switch domain {
		case "djinni":
			c.OnHTML(".list-jobs__title .profile", handleHTML)
		case "dou":
			c.OnHTML(".vacancy .title .vt", handleHTML)
		}

		c.Visit(url)
	}

	return vacanciesSlice
}
