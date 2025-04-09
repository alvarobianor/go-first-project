package utils

import (
	"errors"
	"fmt"
	"math"
)

func Divider(numerator float64, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, errors.New("denominator is 0")
	}

	return numerator / denominator, nil
}

type UserAlvim struct {
	Name string
	Age  int
}

func (u *UserAlvim) ChangeName(name string) {
	u.Name = name
}

func (u *UserAlvim) ChangeAge(age int) {
	u.Age = age
}

func (u *UserAlvim) ItsMe() {
	fmt.Println("My name is", u.Name, "and I am", u.Age, "years old")
}

func CreateUserAlvim(name string, age int) (*UserAlvim, error) {
	if age > 100 {
		return nil, errors.New("age is greater than 100")
	}

	return &UserAlvim{Name: name, Age: age}, nil
}

type ErrorAlvim struct {
	msg string
}
type Error interface {
	Error() string
} // not used

func (e ErrorAlvim) Error() string {
	return e.msg
}

func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrorAlvim{msg: "x is less than 0"}
	}

	return math.Sqrt(x), nil
}
