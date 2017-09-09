package main

import (
	"bufio"
	"math/rand"
	"os"
)

var wordsList []string
var totalNumberOfWords int

func init() {
	wordsList, _ = readWordsList()
	totalNumberOfWords = len(wordsList)
}

// GenerateNextWord picks up a random a word from the list of words
func GenerateNextWord() string {
	return wordsList[rand.Intn(totalNumberOfWords)] + " "
}

func readWordsList() ([]string, error) {
	file, err := os.Open("words.list")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
