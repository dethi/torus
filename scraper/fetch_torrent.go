package scraper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dethi/torus/util"

	"golang.org/x/net/html"
)

func FetchTorrent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("FetchTorrent: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("FetchTorrent: invalid URL: %v", resp.Status)
	}

	if !strings.HasSuffix(url, ".torrent") {
		links, err := extractLinks(resp)
		if err != nil {
			return nil, fmt.Errorf("FetchTorrent: %v", err)
		}

		links = util.Filter(links, func(e string) bool {
			return strings.HasSuffix(e, ".torrent")
		})

		if c := len(links); c != 1 {
			return nil, fmt.Errorf("FetchTorrent: found %v torrents", c)
		}
		return FetchTorrent(links[0])
	}

	if resp.Header.Get("Content-Type") != "application/x-bittorrent" {
		return nil, fmt.Errorf("FetchTorrent: invalid Content-Type: %v",
			resp.Header.Get("Content-Type"))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("FetchTorrent: %v", err)
	}
	return data, err
}

func extractLinks(resp *http.Response) ([]string, error) {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("extractTorrent: parsing HTML: %v", err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
