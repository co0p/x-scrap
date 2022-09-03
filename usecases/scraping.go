package usecases

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

type Scraper interface {
	Scrape(url *url.URL) ([]string, error)
}

type ScrapingCmd struct {
	Url  string
	Tags []string
}

type ScrapingResult struct {
	Url     string
	Tags    []string
	Matches map[string]int
}

type Scraping struct {
	Scraper Scraper
}

func (s Scraping) Scrape(cmd ScrapingCmd) (ScrapingResult, error) {

	u, err := url.Parse(cmd.Url)
	if err != nil {
		return ScrapingResult{}, errors.New("given url is not parsable: " + err.Error())
	}

	paragraphs, err := s.Scraper.Scrape(u)
	if err != nil {
		return ScrapingResult{}, errors.New("failed scraping: " + err.Error())
	}

	rgx := regexp.MustCompile(strings.Join(cmd.Tags, "|"))
	matches := make(map[string]int)
	for _, p := range paragraphs {
		found := rgx.FindAllString(p, -1)

		for _, match := range found {
			matches[match] = matches[match] + 1
		}
	}

	return ScrapingResult{
		Url:     cmd.Url,
		Tags:    cmd.Tags,
		Matches: matches,
	}, nil
}
