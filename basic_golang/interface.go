package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animals struct {
	Name string
}

func Speak(value HasName) {
	fmt.Println("Yo ! ", value.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animals) GetName() string {
	return animal.Name
}

// we also have empty interface which is...
func emptyInterface() any {
	return "empty"
}

func main() {
	person := Person{Name: "Yoku"}
	animal := Animals{Name: "Yeti"}
	Speak(person)
	Speak(animal)

	var zero any = emptyInterface() // in empty inteface it will implement every data types that passed
	fmt.Println(zero)
}
