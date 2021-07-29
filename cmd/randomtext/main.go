package main

import (
	"flag"
	"fmt"
	"github.com/kishaningithub/randomtext"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sizePtr := flag.String("size", "1MB", "Size of generated random text in KB, MB, GB, TB")
	typePtr := flag.String("type", "chars", "Type of text to be generated - chars, words, zeros")
	flag.Parse()
	sizeInBytes := parseSize(*sizePtr)
	generator := parseType(*typePtr)
	randomtext.Generate(sizeInBytes, generator, os.Stdout)
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
		sizeInBytes = givenSize * 1024 * 1024 * 1024 * 1024
	}
	return sizeInBytes
}

func parseType(typeStr string) func() string {
	switch typeStr {
	case "chars":
		return randomtext.CharGenerator()
	case "words":
		return randomtext.WordGenerator()
	case "zeros":
		return randomtext.ZeroGenerator()
	default:
		fmt.Print("Not a valid value")
		os.Exit(1)
	}
	return nil
}
