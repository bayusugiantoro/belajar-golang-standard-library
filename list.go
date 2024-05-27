package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Bayu")
	data.PushBack("Sugiantoro")
	data.PushBack("Bagus")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Bayu

	next := head.Next() // Sugiantoro
	fmt.Println(next.Value)

	next = next.Next() // Bagus
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}