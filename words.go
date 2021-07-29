package randomtext

import (
	_ "embed"
	"strings"
)

//go:embed assets/words.list
var Words string

// WordGenerator returns random words
func WordGenerator() func() string {
	wordsList := strings.Split(Words, "\n")
	return randomGenerator(wordsList)
}
