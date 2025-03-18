package utils

import "math/rand"

// GetRandomNumber retorna um número aleatório
func GetRandomNumberAlv() float64 {
    return generateNumber()
}

func generateNumber() float64 {
    var result float64 = (rand.Float64() + rand.Float64())*3
    return result
}