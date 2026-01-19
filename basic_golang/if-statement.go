package main

import "fmt"

func main() {
	var name = "Yokunoshita"
	// just like another if statement, but we can also do this here:
	if length := len(name); length > 10 {
		fmt.Println("What the hell")
	} else {
		fmt.Println("Wzupp")
	}
}
