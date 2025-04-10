package utils

import (
	"fmt"
	"io"
	"strings"
)

type MyWriter struct{}

func (w MyWriter) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}

func ReadStringAsStrem(str string) {
	reader := strings.NewReader(str)
	myWriter := MyWriter{}

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
