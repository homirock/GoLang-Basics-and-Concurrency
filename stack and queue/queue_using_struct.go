package main

import "fmt"

type queue struct {
	elements []int
	size     int
}

func (m *queue) push(element int) {
	if len((*m).elements) == ((*m).size) {
		return
	}
	(*m).elements = append((*m).elements, element)

}
func (m *queue) pop() {
	if len((*m).elements) == 0 {
		return
	}
	(*m).elements = (*m).elements[1:]
}

func main() {
	q := queue{size: 3}
	q.push(1)
	fmt.Println(q.elements)
	q.push(2)
	fmt.Println(q.elements)
	q.push(3)
	fmt.Println(q.elements)
	q.pop()
	fmt.Println(q.elements)
}
