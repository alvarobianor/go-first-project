package utils

import (
	"io"
	"os"
	"strings"
)

type MyWriter struct {
	reader io.Writer
}

func (w MyWriter) Write(b []byte) (n int, err error) {

	for i, bytes := range b {
		b[i] = bytes + 10
	}

	return w.reader.Write(b)
}

func ReadStringAsStrem(str string) {
	reader := strings.NewReader(str)
	myWriter := MyWriter{reader: os.Stdout}

	buffer := make([]byte, 2)

	for {
		len, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		_, _ = myWriter.Write([]byte(buffer[:len]))

	}

}
