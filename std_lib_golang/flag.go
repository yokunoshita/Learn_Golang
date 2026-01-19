package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Input your host name")
	username := flag.String("username", "root", "Input your username")
	password := flag.String("password", "root", "Input your password")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}
