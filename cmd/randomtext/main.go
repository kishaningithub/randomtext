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
	sizeInBytes := parseSize(strings.ToLower(*sizePtr))
	generator := parseType(strings.ToLower(*typePtr))
	randomtext.Generate(sizeInBytes, generator, os.Stdout)
}

func parseSize(size string) int {
	reg := regexp.MustCompile("\\d+")
	givenSize, _ := strconv.Atoi(reg.FindString(size))
	sizeInBytes := 1
	if strings.Contains(size, "kb") {
		sizeInBytes = givenSize * 1024
	} else if strings.Contains(size, "mb") {
		sizeInBytes = givenSize * 1024 * 1024
	} else if strings.Contains(size, "gb") {
		sizeInBytes = givenSize * 1024 * 1024 * 1024
	} else if strings.Contains(size, "tb") {
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
