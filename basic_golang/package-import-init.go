package main

import (
	_ "basic_golang/blank" // this is how blank identifier is called
	"basic_golang/database"
	"basic_golang/helper"
	"fmt"
)

func main() {
	res := helper.Jello("Yoku")
	fmt.Println(res)

	appVer := helper.Application
	fmt.Println(appVer)

	appNew := helper.TryOne()
	fmt.Println(appNew)

	// initialize func when a package is called
	fmt.Println(database.GetDatabase())

}
