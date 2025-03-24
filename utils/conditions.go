package utils

func Conditional(value uint8) {
	if x := value; x > 5 {
		println("x is greater than 5 ->", x)
	} else if x == 5 {
		println("x is equal to 5 ->", x)
	} else {
		println("x is less than 5 ->", x)
	}
}

func WhatTypeIsThis(value any) {
	switch value.(type) {
	case int:
		println("This is a int")
	case string:
		println("This is a string")
	case float64:
		println("This is a float64")
	default:
		println("This is a unknown type")
	}
}
