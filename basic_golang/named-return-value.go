package main

import "fmt"

// the return value's name is costumized
func getFullName() (Firstname, Lastname string) {
	Firstname = "Yuko"
	Lastname = "Yokubo"

	return Firstname, Lastname
}

func main() {
	firstname, lastname := getFullName()
	fmt.Println(firstname, lastname)
}
