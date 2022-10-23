package usecases

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

type Scraper interface {
	Paragraphs(url *url.URL) ([]string, error)
}

type FindTagsCmd struct {
	Url  string
	Tags []string
}

type FindTagsResult struct {
	Url     string
	Tags    []string
	Matches map[string]int
}

type FindTagsUsecase struct {
	Scraper Scraper
}

func (s FindTagsUsecase) FindTags(cmd FindTagsCmd) (FindTagsResult, error) {

	u, err := url.Parse(cmd.Url)
	if err != nil {
		return FindTagsResult{}, errors.New("given url is not parsable: " + err.Error())
	}

	paragraphs, err := s.Scraper.Paragraphs(u)
	if err != nil {
		return FindTagsResult{}, errors.New("failed extracting paragraphs: " + err.Error())
	}

	rgx := regexp.MustCompile("(?i)" + strings.Join(cmd.Tags, "|"))
	matches := make(map[string]int)
	for _, p := range paragraphs {
		found := rgx.FindAllString(p, -1)
		for _, match := range found {
			tag := strings.ToLower(match)
			matches[tag]++
		}
	}

	return FindTagsResult{
		Url:     cmd.Url,
		Tags:    cmd.Tags,
		Matches: matches,
	}, nil
}
