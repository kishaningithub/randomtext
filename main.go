package randomtext

import (
	"bufio"
	"os"
)

func generate(generator func() string, totalSize int) {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	bytesSent := 0
	for bytesSent < totalSize {
		content := generator()
		noOfBytesWritten, _ := f.WriteString(content)
		if noOfBytesWritten%100 == 0 {
			f.WriteString("\n")
		}
		bytesSent += noOfBytesWritten
	}
}
