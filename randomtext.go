package randomtext

import (
	"bufio"
	"io"
	"math/rand"
)

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go assets/

// Generate text of given type for a given size and writes it to the given writer
func Generate(sizeInBytes int, generate func() string, writer io.Writer) {
	bufferedWriter := bufio.NewWriter(writer)
	defer bufferedWriter.Flush()
	bytesSent := 0
	for bytesSent < sizeInBytes {
		content := generate()
		noOfBytesWritten, _ := bufferedWriter.WriteString(content)
		bytesSent += noOfBytesWritten
	}
}

func randomGenerator(sourceData []string) func() string {
	sourceDataSize := len(sourceData)
	randomArrayIndexes := rand.Perm(sourceDataSize)
	i := -1
	return func() string {
		i = (i + 1) % sourceDataSize
		return sourceData[randomArrayIndexes[i]]
	}
}
