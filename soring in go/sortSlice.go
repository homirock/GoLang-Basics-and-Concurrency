// Go program to illustrate
// how to sort a slice
package main

import (
	"fmt"
	"sort"
)

// Main function
func main() {

	// Creating and initializing
	// a structure
	Author := []struct {
		a_name string
		a_article int
		a_id	 int
	}{
		{"Mina", 304, 1098},
		{"Cina", 634, 102},
		{"Tina", 104, 105},
		{"Rina", 10, 108},
		{"Sina", 234, 103},
		{"Vina", 237, 106},
		{"Rohit", 56, 107},
		{"Mohit", 300, 104},
		{"Riya", 4, 101},
		{"Sohit", 20, 110},
	}

	// Sorting Author by their name
	// Using Slice() function
	sort.Slice(Author, func(p, q int) bool {
		return Author[p].a_name < Author[q].a_name
	})
	
	fmt.Println("Sort Author according to their names:")
	fmt.Println(Author)

	
}
