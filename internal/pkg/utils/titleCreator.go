package utils

import (
	"strconv"
	"strings"
)

func TitleCreator(i int, str string) string {
	strSplit := strings.Split(str, " ")

	title := strconv.Itoa(i)

	exceptionChar := "+=[]:;«,./?|-()\\'"

	for _, s := range strSplit {
		if s == "|" {
			continue
		} else if s == "" {
			continue
		}
		for i := 0; i < len(exceptionChar); i++ {
			s = strings.ReplaceAll(s, string(exceptionChar[i]), "")
		}

		title += "_" + s
	}

	title += ".md"

	return title
}
