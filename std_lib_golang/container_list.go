package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Yoku")         // Back
	data.PushFront("Yokunoshita") // Front
	data.PushBack("Melancholous") // Another Back

	var head list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next().Value
	fmt.Println(next)

	next :=
		fmt.Println(next)
}
