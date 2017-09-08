package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+,./;'[]<>?:{}"

func main() {
	sizePtr := flag.String("size", "1MB", "Size of generated random text")
	flag.Parse()
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
	bytesSent := 0
	for bytesSent < sizeInBytes {
		randomChar := string(charset[rand.Intn(89)])
		fmt.Print(randomChar)
		bytesSent += len(randomChar)
	}
}
