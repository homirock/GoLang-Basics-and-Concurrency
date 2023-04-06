package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func main() {
	list := LinkedList{nil, 0}

	fmt.Println(list)
}
