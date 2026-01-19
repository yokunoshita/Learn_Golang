package main

import "fmt"

type Employee struct {
	name, address string
	age           int
}

func (employee Employee) speak(name string) {
	fmt.Println("It's I ", employee.name, " speak to ", name)
}

func main() {
	var person1 Employee
	person1.age = 20
	person1.name = "this shi"
	person1.address = "isekai"

	fmt.Println(person1)

	// we can also do this to create a struct
	// named struct literals

	person2 := Employee{
		name:    "Yoku",
		address: "Not this world",
		age:     20,
	}
	fmt.Println(person2)

	// or or, we can do this instead
	person3 := Employee{"Budi", "Another we go", 40}
	fmt.Println(person3)

	person1.speak("Mother")
	person2.speak("needa")
}
