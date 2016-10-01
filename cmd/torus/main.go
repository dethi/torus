package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/dethi/torus/web"
)

var version string

func main() {
	configPath := flag.String("config", "torus.cfg", "config pathname")
	versionFlag := flag.Bool("v", false, "print torus version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("rev", version)
		return
	}

	if err := LoadConfig(*configPath); err != nil {
		log.Fatal(err)
	}

	setup()
}

func setup() {
	log.Printf("Serving %v", cfg.ListenAddr)
	log.Fatal(http.ListenAndServe(cfg.ListenAddr, nil))
}
