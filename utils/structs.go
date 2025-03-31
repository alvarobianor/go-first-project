package utils

type MyString string

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	Course string
}

func (p *Person) MyBirthday() {
	p.Age++
}

func (s *Student) ChangeName(name string) {
	s.Name = name
}

func ChangeName2(s *Student, name string) {
	s.Name = name
}

func CreateStudent(name string, age int, course string) Student {
	return Student{
		Person: Person{
			Name: name,
			Age:  age,
		},
		Course: course,
	}
}
