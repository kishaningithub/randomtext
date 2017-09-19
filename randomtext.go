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
	return func() string {
		return sourceData[rand.Intn(sourceDataSize)]
	}
}
