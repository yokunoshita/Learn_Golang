package main

import "fmt"

// variadic func create an array to catch multiple params that will be passed on the functions
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
func main() {
	// this is how variadic func is used
	theValue := sumAll(100, 80, 1987, 2002, 999)
	fmt.Println("sum of all = ", theValue)

	// if we have the data that already been in a slice data type
	slice := []int{10, 20, 30, 40}
	sliceSum := sumAll(slice...)
	fmt.Println("The slice sum value = ", sliceSum)
}
