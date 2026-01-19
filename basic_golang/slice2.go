package main

import "fmt"

func main() {
	// here we about to create slice using make() that will also create an array for it
	oldSlice := make([]string, 2, 5)
	oldSlice[0] = "hello w"
	oldSlice[1] = "Orld"
	// oldSlice[2] = "Za Warudo" will create an error because we declare about to make 2 value of slice.
	// we can add another value using append()

	fmt.Println(oldSlice)

	// Copy slice
	newestSlice := make([]string, len(oldSlice), cap(oldSlice)) // creating new slice based on latest slice
	copy(newestSlice, oldSlice)

	fmt.Println(newestSlice)
}
