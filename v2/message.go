package error_handler

import "strings"

func setErrorMessage(errString string) string {
	errorMessage := ""
	switch {
	case strings.Contains(errString, ErrNilPointer):
		errorMessage = "You are trying to reference a null element"
	case strings.Contains(errString, ErrOutOfRange):
		errorMessage = "An array element is out of range"
	case strings.Contains(errString, ErrDivideByZero):
		errorMessage = "You cannot divide by zero"
	default:
		errorMessage = errString
	}
	return errorMessage
}
