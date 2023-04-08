package main

import "fmt"

type Employee struct {
	name string
	id   int
}

//we can update the properties of struct by passing pointer as receiver
func (a *Employee) UpdateWithPointer(nm string) {
	a.name = nm
}

//we can't update the properties of struct by passing value as receiver
func (a Employee) UpdateWithoutPointer(nm string) {
	a.name = nm
}

func main() {

	a := Employee{name: "partha", id: 2}
	fmt.Println("before Updating name:",a)
	a.UpdateWithPointer("homi")
	fmt.Println("After Updating the name:",a)
	b:=Employee{name: "soumya", id:3}
	fmt.Println("before Updating name:",b)
	// go will undestand .no need to reference the struct.
	// (*b).UpdateWithoutPointer("sidhhi")
	b.UpdateWithoutPointer("sidhhi")
	
	fmt.Println("After Updating name:",b)

}