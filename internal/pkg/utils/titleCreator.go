package utils

import (
	"strconv"
	"strings"
)

func TitleCreator(i int, str string) string {
	strSplit := strings.Split(str, " ")

	title := strconv.Itoa(i)

	for _, s := range strSplit {
		s = stringDeleteExtraChar(s)

		if s == "" {
			continue
		}
		title += "_" + s
	}

	title += ".md"

	return title
}

func stringDeleteExtraChar(s string) string {
	exceptionChar := "+=[]:;«,./? |_-()@#'"
	for i := 0; i < len(exceptionChar); i++ {
		s = strings.ReplaceAll(s, string(exceptionChar[i]), "")
	}
	return s
}
