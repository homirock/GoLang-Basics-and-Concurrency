// Write a function that reverses a string. The input string is given as an array of characters s.

// You must do this by modifying the input array in-place with O(1) extra memory.

// Example 1:

// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]

package main

import "fmt"

func reverseString(s []string) []string {

	size := len(s)
	fmt.Println(size)

	// reverse string by mirror image
	for i := 0; i < size/2; i++ {
		s[i], s[size-1-i] = s[size-1-i], s[i]
	}

	return s

}

func main() {
	a := []string{"a","b"}
	
	fmt.Println(reverseString(a))
}
