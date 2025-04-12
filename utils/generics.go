package utils

import (
	"fmt"
)

// NumberComparer is an interface for types that can be compared
type NumberComparer[T int | float64] interface {
	Max(a, b T) T
}

// NumberComparerImpl implements the NumberComparer interface
type NumberComparerImpl[T int | float64] struct {
	whatTheType string
}

// Max implements the Max function of the NumberComparer interface
func (nc NumberComparerImpl[T]) Max(a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Example usage of NumberComparer
func ExampleNumberComparer() {
	numberComparer := NumberComparerImpl[int]{
		whatTheType: "just for testing",
	}

	max := numberComparer.Max(983, 2983)

	fmt.Println(max)
}

// StackInterface defines the interface for stack operations
type StackInterface[T any] interface {
	Push(item T)
	Pop() (T, bool)
	Peek() (T, bool)
	Size() int
}

// Stack is a generic implementation of a stack
type Stack[T any] struct {
	items []T
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the element from the top of the stack
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek returns the element from the top without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Size returns the current size of the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

func ExampleStack() {
	stack := Stack[float32]{}

	stack.Push(3.14)
	stack.Push(6.28)

	fmt.Println(stack.items)

	stack.Pop()

	fmt.Println(stack.items)

}

type CustomArrays interface {
	int | string
}

type MyType[T CustomArrays] interface {
	Add(item T)
	Show()
	Get() *[]T
}

type MyStruct[T CustomArrays] struct {
	Value []T
}

func (s *MyStruct[T]) Add(item T) {
	s.Value = append(s.Value, item)
}

func (s *MyStruct[T]) Show() {
	fmt.Println(s.Value)
}
func (s *MyStruct[T]) Get() *[]T {
	return &s.Value
}
func ExampleMyStruct() {
	myStruct := MyStruct[int]{Value: []int{1, 2, 3}}

	myStruct.Show()

	value := myStruct.Get()

	*value = append(*value, 4)

	fmt.Println(*value)
	myStruct.Add(5)
	myStruct.Show()

}

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Saudacao() string {
	return fmt.Sprintf("Olá, meu nome é %s", p.Nome)
}

func (p *Pessoa) Aniversario() {
	p.Idade++
}

func ExamplePessoa() {
	alvaro := Pessoa{
		Nome:  "Alvaro",
		Idade: 20,
	}

	alvaro.Saudacao()

	alvaro.Aniversario()

	fmt.Println(alvaro)
}
