package main

import "fmt"

type Man struct {
	name string
}

func (man *Man) Married() {
	man.name = "Mr. " + man.name
}

func (man *Man) UnMarried() {
	man.name = "Bro " + man.name
}

func main() {
	Yoku := Man{"Yoku"}

	Yoku.UnMarried()

	fmt.Println(Yoku.name)
}
