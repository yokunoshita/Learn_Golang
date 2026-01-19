package main

import "fmt"

type Blacklist func(string) bool

func blacklistedName(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You have been blocked")
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	// instead of creating a new block code of function, we declare the function on main func
	blacked := func(name string) bool {
		return name == "anjing"
	}

	blacklistedName("Yoku", blacked) // then we called the function here
	blacklistedName("babi", func(name string) bool {
		return name == "babi"
	}) // we can also create it like this

}
