package main

import "fmt"

//LIFO
// push,pop,check size of stack,get all the item present in the stack,check stack is empty or not,
type Elements []string

func (s *Elements) push(str string) Elements {
	*s = append(*s, str)
	return *s
}

func (s *Elements) pop(str string) Elements {
	*s=(*s)[:len(*s)-1]
	return *s
}

func main() {
	e := Elements{}
	fmt.Println(e.push("x"))
	fmt.Println(e.push("y"))
	fmt.Println(e.push("z"))
	fmt.Println(e.pop("z"))


}
