package scraper

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dethi/torus/util"
	"github.com/pkg/errors"
)

func ScrapeTorrentURL(url string) (string, error) {
	if strings.HasSuffix(url, ".torrent") {
		return url, nil
	}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}

	var res []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if v, ok := s.Attr("href"); ok && strings.HasSuffix(v, ".torrent") {
			res = append(res, v)
		}
	})

	if len(res) != 1 {
		return "", errors.Errorf("found %v torrent URL", len(res))
	}

	return util.AbsoluteURL(url, res[0])
}

func FetchTorrent(url string) ([]byte, error) {
	torrentURL, err := ScrapeTorrentURL(url)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(torrentURL)
	if err != nil {
		return nil, errors.Wrapf(err, "fetch torrent failed %v", torrentURL)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("got `%v` for %v", resp.Status, torrentURL)
	}

	if resp.Header.Get("Content-Type") != "application/x-bittorrent" {
		return nil, errors.Errorf("got `Content-Type: %v` for %v",
			resp.Header.Get("Content-Type"), torrentURL)
	}

	return ioutil.ReadAll(resp.Body)
}
