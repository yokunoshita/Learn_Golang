package main

import "fmt"

func main() {
	// we can create a map using two ways
	person := map[string]string{
		"name":    "Yoku",
		"address": "newWorld",
	}

	fmt.Println(person)

	books := make(map[string]string)
	books["title"] = "Basic Golang"
	books["years"] = "2045"
	books["author"] = "This-shi"
	books["trash"] = "about to be deleted"

	fmt.Println(books)

	// to delete the value on a map
	delete(books, "trash")

	fmt.Println(books)
}
