package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		channel <- "Yoku"
		fmt.Println("Done transfering goroutine to channel")
	}()

	data := <-channel

	fmt.Println("Data on channel : ", data)
	// close(channel)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Yokudake"
}

func TestChannelAsParams(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println("The data that passed by channel on param is : ", data)
	close(channel)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "It's Yoku boi"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("Data from channel : ", data)
}

func TestIOChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}
