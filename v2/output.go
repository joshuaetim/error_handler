package error_handler

import (
	"fmt"
	"strings"
	"unicode"
)

func generateUnderlineWithPrecedingText(fileContents []string, color Color) (string, string) {
	// the line with the error
	// "preceding" is the non-letter character to be added to the "underline"
	// "countBefore" counts spaces to skip before the "underline"
	// color of underline is also set

	preceding, countBefore := "", 0
	codeTextOfInterest := fileContents[5]
	for i := 0; i < len(codeTextOfInterest); i++ {
		if !unicode.IsLetter(rune(codeTextOfInterest[i])) {
			countBefore++
			preceding = fmt.Sprintf("%s%c", preceding, codeTextOfInterest[i])
		} else {
			break
		}
	}
	lineUnder := strings.Repeat("~", len(codeTextOfInterest)-countBefore)
	lineUnder = fmt.Sprintf("%s%s%s", string(color), lineUnder, ColorReset)

	return preceding, lineUnder
}

func printCodeTrace(fileContents []string, lineInt int64, preceding, lineUnder string) {
	// print each line
	for i := 0; i < len(fileContents); i++ {
		lineNum := int(lineInt) + i + -5
		fmt.Printf("%d  %s\n", lineNum, fileContents[i])
		if lineNum == int(lineInt) {
			fmt.Printf("%s%s\n", preceding, lineUnder)
		}
	}
}
