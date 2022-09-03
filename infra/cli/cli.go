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

func (cli *CLI) Execute(args []string) (usecases.ScrapingCmd, error) {

	flags := flag.NewFlagSet("", flag.ContinueOnError)
	url := flags.String("url", "", "the url to scrape")
	taglist := flags.String("tags", "", "comma seperated list of tags")

	flags.Parse(args[1:])

	if len(*url) == 0 {
		return usecases.ScrapingCmd{}, errors.New("flag '-url' must not be empty")
	}

	if len(*taglist) == 0 {
		return usecases.ScrapingCmd{}, errors.New("flag '-tags' must not be empty")
	}

	return usecases.ScrapingCmd{
		Url:  *url,
		Tags: strings.Split(*taglist, ","),
	}, nil

}
