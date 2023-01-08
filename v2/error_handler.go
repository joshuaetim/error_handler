package error_handler

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
)

type Config struct {
	PreferredColor Color
}

type ErrorHandler struct {
	Config *Config
}

// NewErrorHandler returns a new instance of ErrorHandler
func NewErrorHandler(config *Config) *ErrorHandler {
	if config == nil {
		config = &Config{}
	}
	return &ErrorHandler{
		Config: config,
	}
}

// HandleError handles the panic and returns error
func (h *ErrorHandler) HandleError() {
	if err := recover(); err != nil {
		recover()
		stackString := string(debug.Stack())
		stackArray := strings.Split(stackString, "\n")

		filename, line, fnName := getMetaDataFromStack(stackArray)

		// error messages
		errString := fmt.Sprintf("%v", err)
		errorMessage := setErrorMessage(errString)

		fmt.Printf("there was an error:\n\tFile: %s\n\tFunction: %s\n\tLine: %s\n\tMessage: %s\n", filename, fnName, line, errorMessage)

		// line break
		fmt.Println(strings.Repeat("=", 100))

		// code source extraction
		firstSlash := strings.Index(filename, "/")

		// get code from file
		lineInt, _ := strconv.ParseInt(line, 10, 64)
		fileContents := extractCodeFromFile(filename[firstSlash:], lineInt)

		preceding, lineUnder := generateUnderlineWithPrecedingText(fileContents, h.Config.PreferredColor)

		printCodeTrace(fileContents, lineInt, preceding, lineUnder)
	}
}
