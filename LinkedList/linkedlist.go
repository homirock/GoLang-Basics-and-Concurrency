package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) insertAtEnd(val int) {
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

func(l *LinkedList) insertAtStart(val int){
    newNode:=&Node{data:val}
    if l.head == nil {
        l.head=newNode
        return
    }else{
        newNode.next = l.head
        l.head=newNode
    }

}

func(l *LinkedList) inserAtMiddle(val int,position int){
    newNode := &Node{data:val,next:nil}
    if position==0 {
        newNode.next = l.head
        l.head=newNode
        return
    }else{
    current:=l.head
    for i:=0;i<position-1;i++{
        current=current.next
    }
    if current!=nil{
        newNode.next = current.next
        current.next = newNode
    }else{
        fmt.Printf("out of bounce")
    }
}
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
    ll.insertAtEnd(10)
    ll.insertAtEnd(20)
    ll.insertAtEnd(30)
    ll.insertAtStart(40)
    ll.inserAtMiddle(5,2)
    ll.display()
}
