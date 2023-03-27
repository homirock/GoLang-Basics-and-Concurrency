// Golang program to illustrate how
// to create data isolation
package main

import "fmt"
//GFG := 0
// newCounter function to
// isolate global variable
func newCounter() func() int {
	GFG := 0
	return func() int {
	GFG += 1
	return GFG
}
}
func main() {
	// newCounter function is
	// assigned to a variable
	counter := newCounter()

	// invoke counter
	fmt.Println(counter())
	// invoke counter
	fmt.Println(counter())
	
	
}
// Explanation: The closure references the variable GFG even after the newCounter() function has finished 
// running but no other code outside of the newCounter() function has access to this variable. 
// This is how data persistency is maintained between function calls while also isolating the data from other 
// code.