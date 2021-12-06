package curl

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

// Get *html.Node from URL
func GetHtmlNode(url string) (*html.Node, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Errorf("URL: %s. Status code error: %d %s", url, res.StatusCode, res.Status)
		return nil, errors.New("Status code error")
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func GetTextfromHtmlNode(n *html.Node) (string, error) {
	if n.Type == html.TextNode {
		return n.Data, nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return GetTextfromHtmlNode(c)
	}

	return "", errors.New("cant get Text")
}
