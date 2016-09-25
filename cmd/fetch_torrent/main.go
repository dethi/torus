package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dethi/torus/scraper"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: ./fetch_torrent FILENAME URL")
		os.Exit(1)
	}

	b, err := scraper.FetchTorrent(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.Write(b); err != nil {
		log.Fatal(err)
	}
}
