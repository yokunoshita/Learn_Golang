package main

import "fmt"

type Address struct {
	Name, address string
}

func main() {
	address1 := Address{"Yoku", "Malibu"}

	address2 := address1

	address2.address = "Mikono"

	fmt.Println(address1)
	fmt.Println(address2)

	// block code above is a prove that go lang is passing variables by value, not reference
	// it means every variable that pass into a method or function or another variable is just a duplicate

	// Meanwhile pointer can create referece to location of data on memmory without duplicate it.
	// In a nutshell with pointer we can pass by reference

	// lets make it
	address3 := &address1
	address3.Name = "Not Yoku"
	fmt.Println(address3)
	fmt.Println(address1)

}
