package main

import "fmt"

type queue []int

func (m *queue) push(element int) []int {
	*m = append(*m, element)
	return *m
}

func (m *queue) pop() []int {
	*m = (*m)[1:]
	return *m
}

func main() {
	n := queue{1,3}
	fmt.Println(n.push(2))
	fmt.Println(n.push(4))
	fmt.Println(n.pop())
}
