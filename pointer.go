package main

import "fmt"

func withoutptr(value int) {
	value = 0
}

func withptr(value *int) {
	*value = 0
}

func main() {
	value := 1

	withoutptr(value)
	fmt.Println("value without ptr:", value)
	withptr(&value)
	fmt.Println("value with ptr:", value)

}
//without ptr cant change value because it is taking the copy of value 
// The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address. Assigning a value to a dereferenced pointer changes the value at the referenced address.
