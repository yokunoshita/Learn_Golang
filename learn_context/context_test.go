package learncontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Creating context
func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithVal(t *testing.T) {
	contextA := context.Background()
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)

}

// Context with cancel
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel()

	fmt.Println(runtime.NumGoroutine())
}

// Context with timeout
func TestContextWithTimeOut(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}

// Context with deadline
func TestContextWithDeadline(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutines : ", runtime.NumGoroutine())
}
