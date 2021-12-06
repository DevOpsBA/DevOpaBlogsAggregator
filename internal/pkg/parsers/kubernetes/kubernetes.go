package kubernetes

import (
	"devOpsBlogsAggregator/internal/pkg/curl"
	"devOpsBlogsAggregator/internal/pkg/parsers"
	"devOpsBlogsAggregator/internal/pkg/utils"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

const (
	BaseURL  = "https://kubernetes.io"
	BlogsURL = "https://kubernetes.io/blog/"

	ArticleURLsFlag = "class"
	ArticleURLsVal  = "td-sidebar-link td-sidebar-link__page "
)

func GetAllArticleURL(url string) []string {
	page, err := curl.GetHtmlNode(url)
	if err != nil {
		log.Fatal(err)
	}

	return parserArticleURLs(page)
}

func parserArticleURLs(n *html.Node) []string {
	var articlesURLs []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == ArticleURLsFlag && attr.Val == ArticleURLsVal {
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						articlesURLs = append(articlesURLs, attr.Val)
					}
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		articlesURLs = append(articlesURLs, parserArticleURLs(c)...)
	}
	return articlesURLs
}

func ParserArticle(URL string, n *html.Node, article *parsers.Article) {
	// Get tags
	if len(article.Tags) == 0 {
		article.Tags = parsers.GetTags("kubernetes")
	}

	// Get URL
	if article.URL == "" {
		article.URL = URL
	}

	if n.Type == html.ElementNode {
		switch n.Data {
		// Get Title
		case "title":
			if article.Title == "" {
				str, _ := curl.GetTextfromHtmlNode(n)
				article.Title = str
			}
		// Get Body
		case "div":
			if article.Body == "" {
				for _, attr := range n.Attr {
					if attr.Key == "class" && attr.Val == "td-content" {
						article.Body = utils.ConvertHTMLToMD(n)
					}
				}
			}
		// Get Create time
		case "time":
			if article.DateCreate.IsZero() {
				time, err := utils.GetTimeFromString(n.FirstChild.Data)
				if err != nil {
					log.Error(err)
				}
				article.DateCreate = time
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ParserArticle(URL, c, article)
	}

	article.DateParse = time.Now()
	// return articlesURLs
}
