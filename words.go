package randomtext

import (
	"strings"
)

// WordGenerator returns random words
func WordGenerator() func() string {
	data, err := Asset("assets/words.list")
	if err != nil {
		panic(err)
	}
	wordsList := strings.Split(string(data), "\n")
	return randomGenerator(wordsList)
}
