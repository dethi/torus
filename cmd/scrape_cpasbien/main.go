package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/dethi/torus/scraper"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./scrape_cpasbien QUERY")
		os.Exit(1)
	}

	res, err := scraper.ScrapeCpasbien(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Name\tSize\tUp\tDown")
	for _, e := range res {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", e.Name, e.Size, e.Up, e.Down)
	}
	w.Flush()
}
