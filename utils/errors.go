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
	Msg string
}
type ErrorCustomByMe interface {
	Error() string
} // not used

func (e ErrorAlvim) Error() string {
	return e.Msg
}

func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrorAlvim{Msg: "x is less than 0"}
	}

	return math.Sqrt(x), nil
}

// errors.Is: Use quando quiser comparar com um erro específico
// errors.As: Use quando quiser verificar se o erro é de um tipo específico e acessar seus métodos/campos

var ErrNotFound = errors.New("não encontrado")

// Example function showing how to use ErrNotFound
func HandleNotFoundError(err error) {
	if errors.Is(err, ErrNotFound) {
		// Handle the specific error
		fmt.Println(err.Error())
	} else {
		fmt.Println("Error is different from ErrNotFound")
	}
}

// Example function showing how to handle a custom error
func HandleCustomError(err error) {
	var customErr ErrorAlvim
	if errors.As(err, &customErr) {
		// the custom now it get the value from err
		fmt.Println(customErr.Msg, err.Error())
	}
}

// Example of error wrapping
func ProcessData(data string) error {
	if len(data) == 0 {
		return fmt.Errorf("process data: %w", errors.New("empty data"))
	}

	// Simulate another operation that might fail
	if len(data) > 100 {
		return fmt.Errorf("process data: %w", &ErrorAlvim{Msg: "data too long"})
	}

	return nil
}

// Example function showing how to handle wrapped errors
func HandleWrappedError(err error) {
	if err == nil {
		return
	}

	// Check if the error contains our custom ErrorAlvim
	var customErr ErrorAlvim
	if errors.As(err, &customErr) {
		fmt.Printf("Found wrapped ErrorAlvim: %v\n", customErr.Msg)
	}

	// Check if the error contains a specific error message
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Found empty data error")
	}

	// Print the full error chain
	fmt.Printf("Full error: %v\n", err)
}
