package main

import "fmt"

func main() {
	increament := 0

	closureFunc := func() {
		fmt.Println("increament - ", increament)
		increament++
	}

	closureFunc()
	closureFunc()
	closureFunc()

	fmt.Println(increament)
}
