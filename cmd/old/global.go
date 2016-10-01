package main

import (
	"flag"
	"regexp"
)

const (
	findUrl = `((?:https?:\/{2})(?:[-\w]+\.)+(?:\w+)(?:\/[-\w]+)*(?:\.[-\w]+)?)`
)

var (
	configPath  = flag.String("config", "torus.cfg", "config pathname")
	versionFlag = flag.Bool("v", false, "prints current version")

	tBucket  = []byte("Torrents")
	listTmpl = MustTemplate("tmpl/list.html")

	regexUrl = regexp.MustCompile(findUrl)
)

var (
	version string
)
