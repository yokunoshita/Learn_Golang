package main

import "fmt"

func sayHello(name string) string {
	return "Wzup " + name
}

func main() {
	hello := sayHello // this shi made function sayHello() became a variable hello
	fmt.Println(hello("Joe"))
}
