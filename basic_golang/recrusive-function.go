package main

import "fmt"

func factorialLoop(value int) int {
	var result int = 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// instead creating function above, we can do this one below

func factorialUsingRecrusive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialUsingRecrusive(value-1) // here we simplify the for loop on function above
	}
}

func main() {
	factor := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(factor)
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialUsingRecrusive(10))
}
