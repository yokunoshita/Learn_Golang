package main

import "fmt"

func main() {
	var name string = "Yoku"

	switch name {
	case "Lord":
		fmt.Println("Not yoku")
	case "yoku":
		fmt.Println("Smoll yoku")
	case "Yoku":
		fmt.Println("The right one")
	default:
		fmt.Println("Wrong name")
	}

	// we can also do short statement on switch
	switch length := len(name); length < 5 {
	case true:
		fmt.Println("Name is too short")
	case false:
		fmt.Println("Name is too long")
	}

	// we can also create if-statement like switch
	var grade int = 80
	var GOOD_GRADE int = 100
	var NOT_BAD_GRADE int = 70
	var BAD_GRADE int = 60

	switch {
	case grade > GOOD_GRADE:
		fmt.Println("The grade is too high")
	case grade > NOT_BAD_GRADE:
		fmt.Println("The grade ain't bad")
	case grade < BAD_GRADE:
		fmt.Println("Go watch the video")
	}
}
