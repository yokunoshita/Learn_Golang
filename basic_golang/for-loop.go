package main

import "fmt"

func main() {
	// default :
	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Loop : ", counter)
	}

	// using range :
	months := []string{"March", "May", "June"}

	for index, month := range months {
		fmt.Println("index", index, "month", month)
	}
	// or if we just wanna return the value without an index :
	for _, month := range months {
		fmt.Println(month)
	}
}
