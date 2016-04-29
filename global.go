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

	algoliaAppId     = flag.String("algoliaAppId", "", "algolia app id")
	algoliaApiKey    = flag.String("algoliaApiKey", "", "algolia api key")
	algoliaSearchKey = flag.String("algoliaSearchKey", "", "algolia search key")
	algoliaIndex     = flag.String("algoliaIndex", "", "algolia index name")

	dbPath       = flag.String("dbPath", "/static/torrents.db", "path to database")
	dataPath     = flag.String("dataPath", "/static/data", "path to data")
	htpasswdPath = flag.String("htpasswdPath", "/static/htpasswd", "path to htpasswd")

	migrate    = flag.Bool("migrate", false, "migrate database")
	activeJobs = flag.Int("activeJobs", 3, "number of active torrents")

	tBucket    = []byte("Torrents")
	listTmpl   = MustTemplate("tmpl/list.html.tmpl")
	regexUrl   = regexp.MustCompile(findUrl)
	regexClean = regexp.MustCompile(cleanName)
)
