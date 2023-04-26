// Given a string s consisting of words and spaces, return the length of the last word in the string.

// A word is a maximal
// substring
//  consisting of non-space characters only.

// Example 1:

// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.
package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	st := strings.TrimSpace(s)
	st1 := strings.Split(st, " ")
	l := len(st1)
	return len(st1[l-1])

}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}