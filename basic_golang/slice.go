package main

import "fmt"

func main() {

	thisIsArray := [...]string{"Yuko", "Yoku", "Yokubo", "Yuno", "Ye"}
	slice := thisIsArray[0:3]
	// can also be written :
	// var slice []string = thisIsArray[0:3]

	fmt.Println(slice[2])
	fmt.Println(slice)

	days := [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	daySlice := days[5:]
	fmt.Println(daySlice)

	// When you change the data on slice based on array,
	// your array will also change it's value based on what you changed on slice
	daySlice[0] = "newDay"
	fmt.Println(daySlice)
	fmt.Println(days)

	// When you append new value on slice, append will check if the array is already full or not.
	// If array is full, then append will create a new array based on the old array + new value
	daySlice2 := append(daySlice, "anotherNewDay")
	fmt.Println(daySlice)  // here will stay the same
	fmt.Println(daySlice2) // here the value from daySlice will be added + it's new value
	fmt.Println(days)      // here will stay the same after changed by daySlice at peverous code block

	// so.. if we about to change some value on daySlice2
	daySlice2[0] = "The day before yesterday"
	// the only one that changed is daySlice2
	fmt.Println(daySlice)
	fmt.Println(daySlice2)
	fmt.Println(days)

}
