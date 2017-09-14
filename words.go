package randomtext

import (
	"fmt"
	"strings"
	"math/rand"
)

// WordGenerator returns random words
func WordGenerator() func() string {
	data, err := Asset("assets/words.list")
	fmt.Println("err", err)
	if err != nil {
		panic(err)
	}
	wordsStr := string(data)
	wordsList := strings.Split(wordsStr, "\n")
	totalNumberOfWords := len(wordsList)
	return func() string {
		return wordsList[rand.Intn(totalNumberOfWords)]
	}
}