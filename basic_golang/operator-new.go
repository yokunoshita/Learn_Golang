package main

import "fmt"

type NewAddress struct {
	city, province, country string
}

func main() {
	var address1 *NewAddress = new(NewAddress)

	address1.city = "Shaw"
	fmt.Println(address1)

	var address2 *NewAddress = address1
	address2.country = "Britain"

	fmt.Println(address1) // it's also changed
	fmt.Println(address2)
}
