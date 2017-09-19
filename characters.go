package randomtext

import (
	"strings"
)

// CharGenerator returns a function that generates random chars
func CharGenerator() func() string {
	data, err := Asset("assets/charecters.list")
	if err != nil {
		panic(err)
	}
	allCharacters := strings.Split(string(data), "\n")
	return randomGenerator(allCharacters)
}
