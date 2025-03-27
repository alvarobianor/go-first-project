package utils

import "math/rand"

func CreateMap() map[string]int {
	return map[string]int{
		"Álvaro": 26,
		"Alvim":  27,
		"Rei":    28,
	}
}

func AddRandoNumberInMap(m map[string]int) {
	m["Random"] = rand.Intn(101)
}

func AddGamesInMap(m map[string][][]uint8, key string, value []uint8) {
	m[key] = append(m[key], value)
}
