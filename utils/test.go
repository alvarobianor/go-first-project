package utils

import (
	"fmt"
	"math/rand"
)

// GetRandomNumber retorna um número aleatório
func GetRandomNumberAlv() float64 {
	return generateNumber()
}

func generateNumber() float64 {
	var result float64 = (rand.Float64() + rand.Float64()) * 3
	return result
}

// A function so funny that can print anything
func PrintAnything(params ...any) []any {
	for index, param := range params {

		fmt.Println("value ->", param)
		fmt.Println("index ->", index)
	}

	return params
}

func ConvertInString(value int) string {
	vale := string(value)

	return vale
}

func TakeInt8(value int8) int8 {
	return value
}
func TakeInt16(value int16) int16 {
	return value
}
