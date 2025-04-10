package utils

import (
	"fmt"
	"io"
	"strings"
)

func ReadStringAsStrem(str string) {
	reader := strings.NewReader(str)

	buffer := make([]byte, 2)

	for {
		len, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(buffer[:len])
	}

}
