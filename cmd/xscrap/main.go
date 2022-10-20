package main

import (
	"fmt"
	"os"

	clipkg "github.com/co0p/x-scrap/infra/cli"
	"github.com/co0p/x-scrap/infra/scraper"
	"github.com/co0p/x-scrap/usecases"
)

func main() {

	cli := clipkg.CLI{}
	colly := scraper.NewColly()
	usecase := usecases.Scraping{Scraper: &colly}

	clicmd, err := cli.Execute(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed parsing input: %s\n", err)
		os.Exit(1)
	}

	for _, url := range clicmd.URLs {

		cmd := usecases.ScrapingCmd{
			Url:  url,
			Tags: clicmd.Tags,
		}

		res, err := usecase.Scrape(cmd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed scraping: %s\n", err)
			os.Exit(1)
		}

		clipkg.Print(os.Stdout, res)
	}
}
