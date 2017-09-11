package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kishaningithub/randomtext"
)

func main() {
	sizePtr := flag.String("size", "1MB", "Size of generated random text")
	typePtr := flag.String("type", "chars", "Type of text to be generated - chars, words")
	flag.Parse()
	sizeInBytes := parseSize(*sizePtr)
	generator := parseType(*typePtr)
	randomtext.Generate(sizeInBytes, generator)
}

func parseSize(sizePtr string) int {
	reg := regexp.MustCompile("\\d+")
	givenSize, _ := strconv.Atoi(reg.FindString(sizePtr))
	sizeInBytes := 1
	if strings.Contains(sizePtr, "KB") {
		sizeInBytes = givenSize * 1024
	} else if strings.Contains(sizePtr, "MB") {
		sizeInBytes = givenSize * 1024 * 1024
	} else if strings.Contains(sizePtr, "GB") {
		sizeInBytes = givenSize * 1024 * 1024 * 1024
	} else if strings.Contains(sizePtr, "TB") {
		sizeInBytes = givenSize * 1024 * 1024 * 1024
	}
	return sizeInBytes
}

func parseType(typeStr string) randomtext.Generator {
	switch typeStr {
	case "chars":
		return randomtext.CharGenerator{}
	case "words":
		return randomtext.WordGenerator{}
	case "zeros":
		return randomtext.ZeroGenerator{}
	default:
		fmt.Print("Not a valid value")
		os.Exit(1)
	}
	return nil
}
