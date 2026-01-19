package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestFuncHelloWorld(t *testing.T) {
	go RunHelloWorld() //this is goroutine
	fmt.Println("Eh...")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 10; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(1 * time.Second)
}
