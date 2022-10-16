package cli

import (
	"errors"
	"flag"
	"strings"

	"github.com/co0p/x-scrap/usecases"
)

type CLI struct{}

func NewCLI() CLI {
	return CLI{}
}

func (cli *CLI) Execute(args []string) ([]usecases.ScrapingCmd, error) {

	flags := flag.NewFlagSet("", flag.ContinueOnError)
	urls := flags.String("url", "", "the urls to scrape (comma <,> separated)")
	taglist := flags.String("tags", "", "comma seperated list of tags")

	flags.Parse(args[1:])

	if len(*urls) == 0 {
		return []usecases.ScrapingCmd{}, errors.New("flag '-url' must not be empty")
	}

	if len(*taglist) == 0 {
		return []usecases.ScrapingCmd{}, errors.New("flag '-tags' must not be empty")
	}

	cmds := extractCommands(urls, taglist)

	return cmds, nil

}

func extractCommands(urls *string, taglist *string) []usecases.ScrapingCmd {
	separatedUrls := strings.Split(*urls, ",")
	separatedTags := strings.Split(*taglist, ",")
	cmds := make([]usecases.ScrapingCmd, len(separatedUrls))

	for i, url := range separatedUrls {
		cmds[i] = usecases.ScrapingCmd{
			Url:  url,
			Tags: separatedTags,
		}
	}
	return cmds
}
