package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sizePtr := flag.String("size", "1MB", "Size of generated random text")
	typePtr := flag.String("type", "chars", "Type of text to be generated - chars, words")
	flag.Parse()
	size := parseSize(sizePtr)
	generator := parseType(typePtr)
	generate(generator, size)
}

func generate(generator func() string, totalSize int) {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	bytesSent := 0
	for bytesSent < totalSize {
		content := generator()
		noOfBytesWritten, _ := f.WriteString(content)
		bytesSent += noOfBytesWritten
	}
}

func parseSize(sizePtr *string) int {
	reg := regexp.MustCompile("\\d+")
	givenSize, _ := strconv.Atoi(reg.FindString(*sizePtr))
	sizeInBytes := 1
	if strings.Contains(*sizePtr, "KB") {
		sizeInBytes = givenSize * 1024
	} else if strings.Contains(*sizePtr, "MB") {
		sizeInBytes = givenSize * 1024 * 1024
	} else if strings.Contains(*sizePtr, "GB") {
		sizeInBytes = givenSize * 1024 * 1024 * 1024
	} else if strings.Contains(*sizePtr, "TB") {
		sizeInBytes = givenSize * 1024 * 1024 * 1024
	}
	return sizeInBytes
}

func parseType(typePtr *string) func() string {
	flag.Parse()
	switch *typePtr {
	case "chars":
		return GenerateNextChar
	case "words":
		return GenerateNextWord
	default:
		fmt.Print("Not a valid value")
		os.Exit(1)
	}
	return nil
}
