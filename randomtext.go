package randomtext

import (
	"bufio"
	"os"
)

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go assets/

// Generate text of given type for a given size
func Generate(sizeInBytes int, generate func() string) {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	bytesSent := 0
	for bytesSent < sizeInBytes {
		content := generate()
		noOfBytesWritten, _ := f.WriteString(content)
		bytesSent += noOfBytesWritten
	}
}
