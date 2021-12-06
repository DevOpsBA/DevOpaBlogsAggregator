package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func TitleCreator(i int, str string) string {
	strSplit := strings.Split(str, " ")

	title := strconv.Itoa(i)

	for _, s := range strSplit {
		if s == "|" {
			continue
		}
		s = stringTrims(s)

		title += "_" + s
	}

	title += ".md"

	return title
}

func stringTrims(s string) string {
	exceptionChar := "+=[]:;«,./? |_-()\\'"
	fmt.Println("Before: ", s)
	for i := 0; i < len(exceptionChar); i++ {
		s = strings.Trim(s, string(exceptionChar[i]))
	}

	fmt.Println("After: ", s)
	return s
}
