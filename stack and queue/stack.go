package main

import "fmt"

//LIFO
// push,pop,check size of stack,get all the item present in the stack,check stack is empty or not,
type Elements []string

func (s *Elements) push(str string) Elements {
	*s = append(*s, str)
	return *s
	//return len(s) == 0
}

func main() {
	e := Elements{}
	fmt.Println(e.push("x"))

}
