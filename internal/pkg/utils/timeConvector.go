package utils

import (
	cErr "devOpsBlogsAggregator/internal/pkg/customErrors"
	"errors"
	"time"
)

func GetTimeFromString(str string) (time.Time, error) {
	var timeFormates = []string{"Monday, January 02, 2006",
		"2006.01.02"}
	var i = 0
	return getTime(str, i, timeFormates)
}

func getTime(str string, i int, timeFormates []string) (time.Time, error) {
	if i >= len(timeFormates) {
		errorMessage := "Cant convert string: " + str + " to time. Need add correct Time Format"
		return time.Now(), &cErr.CustomError{
			PackageName: "utils",
			Functions:   "GetTimeFromString",
			Message:     errorMessage,
			Err:         errors.New("Index out of range")}
	}

	t, err := time.Parse(timeFormates[i], str)
	if err != nil {
		return getTime(str, i+1, timeFormates)
	}

	return t, nil
}
