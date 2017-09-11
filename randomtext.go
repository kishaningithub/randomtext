package randomtext

import (
	"bufio"
	"os"
)

// Generator generates randome text text
//
// Implementations can generage subsets of random text
type Generator interface {
	Generate() string
}

// Generate text of given type for a given size
func Generate(sizeInBytes int, generator Generator) {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	bytesSent := 0
	for bytesSent < sizeInBytes {
		content := generator.Generate()
		noOfBytesWritten, _ := f.WriteString(content)
		bytesSent += noOfBytesWritten
	}
}
