package main

import "fmt"

func deferOne() {
	fmt.Println("This shi will called defer")
}

func callingDefer() {
	defer deferOne() // this shi will be called after callingDeffer() finished running
	fmt.Println("Calling defer in this function")
}

func main() {
	callingDefer()
}
