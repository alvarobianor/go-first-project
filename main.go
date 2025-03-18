package main

import (
	"fmt"
	"myFirstProject/utils"
)

func main() {
	fmt.Println("My favorite number is", utils.GetRandomNumberAlv())
	fmt.Println("Print All things", PrintAnything(1, 3, 4,5, 6,64, 32,3 ,1, "as", '1'))
}

func PrintAnything(params ...any) []any {
	for _, param := range params {
		fmt.Println(param)
	}

	return params
}