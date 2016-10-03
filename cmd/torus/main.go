package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "expvar"
	_ "net/http/pprof"

	"github.com/dethi/torus"
	_ "github.com/dethi/torus/api"
	_ "github.com/dethi/torus/web"
)

func main() {
	configPath := flag.String("config", "torus.cfg", "config pathname")
	versionFlag := flag.Bool("v", false, "print torus version")
	flag.Parse()

	if *versionFlag {
		fmt.Println(torus.Revision())
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
