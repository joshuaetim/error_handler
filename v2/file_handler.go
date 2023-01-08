package error_handler

import (
	"bufio"
	"log"
	"os"
)

func extractCodeFromFile(fname string, lineInt int64) []string {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewScanner(f)

	// extract lines, 5 before, 2 after
	codeText := []string{}
	var lineNum int = 0
	for in.Scan() {
		lineNum++
		if lineNum >= int(lineInt)-5 && lineNum <= int(lineInt)+2 {
			codeText = append(codeText, in.Text())
		}
		if lineNum > int(lineInt)+2 {
			break
		}
	}
	return codeText
}
