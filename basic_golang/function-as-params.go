package main

import "fmt"

// func getNameFiltered(name string, filter func(string) string) {
// 	fmt.Println("Yo " + filter(name))
// }

// the code below is waay more clean than above
// using type declaration Filter

type Filter func(string) string

func getNameFiltered(name string, filter Filter) {
	fmt.Println("Yo " + filter(name))
}

func filterName(name string) string {
	if name == "Babi" {
		return "censored"
	} else {
		return name
	}
}

func main() {
	badwordFilter := filterName

	getNameFiltered("Babi", badwordFilter)
	getNameFiltered("Yoku", badwordFilter)
}
