package parsers

import (
	"time"
)

type Article struct {
	URL        string
	Tags       []string
	Title      string
	Body       string
	DateCreate time.Time
	DateParse  time.Time
}

func GetTags(str string) []string {
	var tags []string

	tags = append(tags, str)
	return tags
}
