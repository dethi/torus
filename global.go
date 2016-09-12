package main

import (
	"flag"
	"regexp"
)

const (
	findUrl   = `((?:https?:\/{2})(?:[-\w]+\.)+(?:\w+)(?:\/[-\w]+)*(?:\.[-\w]+)?)`
	cleanName = `(?i)((\[ *)?[a-z]+.cpasbien.[a-z]+( *\])?)|(web(-?dl)?)|(xvid)`
)

var (
	mgDomain    = flag.String("mgDomain", "", "mailgun domain")
	mgKey       = flag.String("mgKey", "", "mailgun key")
	mgPublicKey = flag.String("mgPublicKey", "", "mailgun public key")

	dbPath       = flag.String("dbPath", "/static/torrents.db", "path to database")
	dataPath     = flag.String("dataPath", "/static/data", "path to data")
	htpasswdPath = flag.String("htpasswdPath", "/static/htpasswd", "path to htpasswd")

	activeJobs = flag.Int("activeJobs", 3, "number of active torrents")

	tBucket    = []byte("Torrents")
	listTmpl   = MustTemplate("tmpl/list.html")
	regexUrl   = regexp.MustCompile(findUrl)
	regexClean = regexp.MustCompile(cleanName)
)
