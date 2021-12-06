package utils

import "golang.org/x/net/html"

func ConvertHTMLToMD(n *html.Node) string {
	var content string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			if c.Data != "" && c.Data != "\n" && c.Data != "\t" {
				content += c.Data
			}
		} else if c.Type == html.ElementNode {
			switch c.Data {
			case "strong":
				content += "**" + ConvertHTMLToMD(c) + "**"
			case "em":
				content += "*" + ConvertHTMLToMD(c) + "*"
			case "a":
				for _, attr := range c.Attr {
					if attr.Key == "href" {
						var hrefText string
						for ch := c.FirstChild; ch != nil; ch = ch.NextSibling {
							hrefText += hrefText + ch.Data
						}
						content += "[" + hrefText + "](" + attr.Val + ")" // [Duck Duck Go](https://duckduckgo.com).
					}
				}
			case "hr":
				content += "\n---\n"
			case "code":
				content += "```" + c.FirstChild.Data + "```"
			case "p":
				content += ConvertHTMLToMD(c) + "\n"
			case "pre":
				content += ConvertHTMLToMD(c)
			}
		}
	}
	return content
}
