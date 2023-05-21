package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) insert(val int) {
	newNode := &Node{data: val, next: nil}

	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}


func (l *LinkedList) display() {
	current := l.head
	for current != nil {
		fmt.Printf("->%v", current.data)
        current=current.next
	}
}

func main(){
    ll:=new(LinkedList)
    ll.insert(10)
    ll.insert(20)
    ll.insert(30)
    ll.display()
}
