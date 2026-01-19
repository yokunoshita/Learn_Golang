package main

import "fmt"

func random() any {
	return "this one story"
}

func main() {
	var result any = random()

	switch value := result.(type) {
	case string:
		fmt.Println("String type :", value)
	case int:
		fmt.Println("Int type :", value)
	default:
		fmt.Println("Unknown data type :", value)
	}
}
