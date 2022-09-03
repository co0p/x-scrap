package main

import (
	"fmt"
	"os"

	clipkg "github.com/co0p/x-scrap/infra/cli"
	"github.com/co0p/x-scrap/infra/scraper"
	"github.com/co0p/x-scrap/usecases"
)

func main() {

	cli := clipkg.NewCLI()
	colly := scraper.NewColly()
	usecase := usecases.Scraping{Scraper: &colly}

	cmd, err := cli.Execute(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed parsing input: %s\n", err)
		os.Exit(1)
	}

	res, err := usecase.Scrape(cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed scraping: %s\n", err)
		os.Exit(1)
	}

	clipkg.Print(os.Stdout, res)
}
