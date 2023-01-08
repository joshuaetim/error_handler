package error_handler

import "strings"

// handles operations related to the stack trace

func getMetaDataFromStack(stackArray []string) (string, string, string) {
	fileNameLine := strings.Split(stackArray[8], ":")
	filename, line := fileNameLine[0], strings.Split(fileNameLine[1], " ")[0]
	fnName := strings.Split(stackArray[7], "(")[0]
	return filename, line, fnName
}
