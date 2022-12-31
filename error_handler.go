package error_handler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"unicode"
)

type Name struct {
	Value string
}

var (
	ErrNilPointer   = "nil pointer dereference"
	ErrOutOfRange   = "index out of range"
	ErrDivideByZero = "divide by zero"
)

func HandleError() {
	if err := recover(); err != nil {
		recover()
		stackString := string(debug.Stack())
		stackArray := strings.Split(stackString, "\n")

		fileNameLine := strings.Split(stackArray[8], ":")
		filename, line := fileNameLine[0], strings.Split(fileNameLine[1], " ")[0]
		fnName := strings.Split(stackArray[7], "(")[0]

		errString := fmt.Sprintf("%v", err)
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
		fmt.Printf("there was an error:\n\tFile: %s\n\tFunction: %s\n\tLine: %s\n\tMessage: %s\n", filename, fnName, line, errorMessage)

		fmt.Println(strings.Repeat("=", 100))
		firstSlash := strings.Index(filename, "/")

		// get code
		f, err := os.Open(filename[firstSlash:])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		in := bufio.NewScanner(f)
		// 5 before, 2 after
		codeText := []string{}
		var lineNum int = 0
		lineInt, _ := strconv.ParseInt(line, 10, 64)
		for in.Scan() {
			lineNum++
			if lineNum >= int(lineInt)-5 && lineNum <= int(lineInt)+2 {
				codeText = append(codeText, in.Text())
			}
			if lineNum > int(lineInt)+2 {
				break
			}
		}
		preceding, countBefore := "", 0

		codeTextOfInterest := codeText[5]
		for i := 0; i < len(codeTextOfInterest); i++ {
			if !unicode.IsLetter(rune(codeTextOfInterest[i])) {
				countBefore++
				preceding = fmt.Sprintf("%s%c", preceding, codeTextOfInterest[i])
			} else {
				break
			}
		}
		lineUnder := strings.Repeat("~", len(codeTextOfInterest)-countBefore)
		// fmt.Printf("%s%s\n", preceding, lineUnder)
		for i := 0; i < len(codeText); i++ {
			lineNum := int(lineInt) + i + -5
			fmt.Printf("%d  %s\n", lineNum, codeText[i])
			if lineNum == int(lineInt) {
				fmt.Printf("%s%s\n", preceding, lineUnder)
			}
		}
	}
}
