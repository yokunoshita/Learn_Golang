package main

import "fmt"

func endCall() {
	fmt.Println("End of service")
	message := recover() // this will resume the program after panic is called
	fmt.Println("this is recover, with the panic : ", message)
}

func callingPanic(error bool) {
	defer endCall() // this defer will be called if panic executed

	if error {
		panic("Panic is called") // this is panic that will terminate the program if the condition is true
	}

	defer endCall() // this defer won't be called after panic is executed
}

func main() {
	callingPanic(true)
}
