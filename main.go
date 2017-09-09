package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	channel := make(chan string)
	defer close(channel)
	sizePtr := flag.String("size", "1MB", "Size of generated random text")
	typePtr := flag.String("type", "chars", "Type of text to be generated - chars, words")
	flag.Parse()
	size := parseSize(sizePtr)
	generator := parseType(typePtr)
	go generateAndSendToChannel(generator, channel)
	channelToStdout(channel, size)
}

func generateAndSendToChannel(generator func() string, channel chan<- string) {
	for {
		channel <- generator()
	}
}

func channelToStdout(reader <-chan string, totalSize int) {
	bytesSent := 0
	for bytesSent < totalSize {
		content := <-reader
		fmt.Print(content)
		bytesSent += len(content)
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
