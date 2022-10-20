package cli

import (
	"errors"
	"flag"
	"strings"
)

type CliCommand struct {
	URLs []string
	Tags []string
}

type CLI struct{}

func (cli *CLI) Execute(args []string) (CliCommand, error) {

	flags := flag.NewFlagSet("", flag.ContinueOnError)
	urls := flags.String("urls", "", "urls to scrape (comma <,> separated)")
	taglist := flags.String("tags", "", "tags to scan for (comma <,> separated)")

	flags.Parse(args[1:])

	if len(*urls) == 0 {
		return CliCommand{}, errors.New("flag '-urls' must not be empty")
	}

	if len(*taglist) == 0 {
		return CliCommand{}, errors.New("flag '-tags' must not be empty")
	}

	cmd := CliCommand{
		URLs: strings.Split(*urls, ","),
		Tags: strings.Split(*taglist, ","),
	}

	return cmd, nil
}
