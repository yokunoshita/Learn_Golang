package main

import (
	"fmt"
	"os"
)

func main() {
	anOs := os.Args

	for _, args := range anOs {
		fmt.Println(args)
	}
}
