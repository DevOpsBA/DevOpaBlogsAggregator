package customErrors

import "fmt"

type error interface {
	Error() string
}

type CustomError struct {
	PackageName string
	Functions   string
	Message     string
	Err         error
}

func (e *CustomError) Error() string {
	errorString := fmt.Sprintf("Package: %s; Function: %s; Message: %s; Error: %s", e.PackageName, e.Functions, e.Message, e.Err.Error())
	return errorString
}
