package scraper

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CpasbienResult struct {
	Name string
	URL  string
	Size string
	Up   int
	Down int
}

const cpasbienURL = "http://www.cpasbien.cm/recherche/"

func ScrapeCpasbien(query string) ([]CpasbienResult, error) {
	resp, err := http.PostForm(cpasbienURL,
		url.Values{"champ_recherche": {query}})
	if err != nil {
		return nil, err
	}

	// GoQuery reads, parses and closes the response body.
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	var res []CpasbienResult
	doc.Find(`div[class^="ligne"]`).Each(func(i int, s *goquery.Selection) {
		titleNode := s.Find(".titre").First()
		sizeNode := s.Find(".poid").First()
		upNode := s.Find(".up").First()
		downNode := s.Find(".down").First()

		res = append(res, CpasbienResult{
			Name: titleNode.Text(),
			URL:  titleNode.AttrOr("href", "#"),
			Size: strings.TrimSpace(sizeNode.Text()),
			Up:   normalizeInt(upNode.Text()),
			Down: normalizeInt(downNode.Text()),
		})
	})

	return res, nil
}

func normalizeInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		n = 0
	}
	return n
}
