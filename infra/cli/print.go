package cli

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/co0p/x-scrap/usecases"
)

func Print(out io.Writer, res usecases.ScrapingResult) {

	fmt.Fprintf(out, "\nurl: %s\n", res.Url)
	fmt.Fprintln(out, "-----------------------")
	w := tabwriter.NewWriter(out, 1, 1, 1, ' ', 0)

	for _, v := range res.Tags {
		c := res.Matches[v]
		fmt.Fprintf(w, "%s:\t%d\n", v, c)
	}
	w.Flush()
}
