package main

import "fmt"

type FullAddress struct {
	city, province, country string
}

func main() {
	address1 := FullAddress{"Wonosobo", "Central Java", "Indonesia"}
	address2 := &address1
	address2.city = "detroit"
	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = FullAddress{"Axcel", "Central Java", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
