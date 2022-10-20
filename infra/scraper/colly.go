package scraper

import (
	"net/url"

	"github.com/gocolly/colly"
)

type Colly struct {
	collector colly.Collector
	found     []string
}

func NewColly() Colly {

	collector := *colly.NewCollector()

	return Colly{
		collector: collector,
	}
}

func (c *Colly) Paragraphs(url *url.URL) ([]string, error) {

	c.restrictDomain(url)
	c.registerParagraphExtractor()
	c.collector.Visit(url.String())

	return c.found, nil
}

func (c *Colly) restrictDomain(url *url.URL) {
	domain := url.Hostname()
	c.collector.AllowedDomains = append(c.collector.AllowedDomains, domain)
}

func (c *Colly) registerParagraphExtractor() {
	c.collector.OnHTML("p", func(e *colly.HTMLElement) {
		c.found = append(c.found, e.Text)
	})
}
