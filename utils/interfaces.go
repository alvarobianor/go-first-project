package utils

import "fmt"

type Animal interface {
	Speak()
	Move()
	GetBone() string
}

func SpeakAnimal(a Animal) {
	a.Speak()
}

func MoveAnimal(a Animal) {
	a.Move()
}

// This Structrepresents a Dog, and implements all methods of the Animal interface, so Go considers it an Animal
// and allows us to use it as an Animal
// and use the methods of the Animal interface
type Dog struct {
	Name string
}

func (d Dog) Speak() {
	println("Woof! My name is", d.Name)
}

func (d Dog) Move() {
	println("I am running!")
}

func (d Dog) GetBone() string {
	return "Bone"
}

func GetBone(a Animal) string {
	return a.GetBone()
}

func (d *Dog) ChangeMyName(name string) {
	d.Name = name
}

func ChangeDogName(dog *Dog, name string) {
	dog.Name = name
}

func CreateDog(name string) Dog {
	return Dog{Name: name}
}

func DemonstrateAnimal() {
	dog := CreateDog("Rex")
	dog.Speak()
	dog.Move()

	SpeakAnimal(dog)
	MoveAnimal(dog)

	ChangeDogName(&dog, "Max")
	dog.Speak()

	dog.ChangeMyName("Buddy")
	dog.Speak()

	bone := "Empty"

	bone = GetBone(dog)

	fmt.Println("Dog's bone:", bone)
}
