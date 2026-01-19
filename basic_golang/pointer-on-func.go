package main

import "fmt"

type Monster struct {
	voice, howToWalk string
}

func AnimalMutation(animal *Monster) {
	animal.voice = "Fahhhhhhh"
}

func main() {
	var animal *Monster = &Monster{}
	AnimalMutation(animal)

	// or or you can also write it :
	animal2 := Monster{}
	AnimalMutation(&animal2)

	fmt.Println(animal)
	fmt.Println(animal2)
}
