package utils

import "fmt"

func DoDefer() int {
	defer fmt.Println("This is the first defer, but will be the fifth")
	defer fmt.Println("This is the second defer, but will be the forth")
	defer fmt.Println("This is the third defer, but will be the third")
	defer fmt.Println("This is the fourth defer, but will be the second")
	defer fmt.Println("This is the fifth defer, but will be the first")
	return 5
}
