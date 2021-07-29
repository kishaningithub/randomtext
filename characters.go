package randomtext

import (
	_ "embed"
	"strings"
)

//go:embed assets/characters.list
var Characters string

// CharGenerator returns a function that generates random chars
func CharGenerator() func() string {
	allCharacters := strings.Split(Characters, "\n")
	return randomGenerator(allCharacters)
}
