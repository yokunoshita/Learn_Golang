package main

import "fmt"

func main() {
	GLOBAL_VAR := 10

	// this is how break is used :
	for i := 1; i < GLOBAL_VAR; i++ {
		if i == 5 {
			fmt.Println("The loop is stopped by break")
			break
		}
		fmt.Println(i)
	}

	// and this is how continue is used :
	for i := 1; i < GLOBAL_VAR; i++ {
		if i%2 == 0 {
			continue // this shi will straight back to loop statement
			// ain't gonna execute the code below
		}
		fmt.Println("Odd numbers = ", i)
	}
}
