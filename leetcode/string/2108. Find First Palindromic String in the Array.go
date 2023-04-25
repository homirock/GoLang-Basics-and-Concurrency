// Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".

// A string is palindromic if it reads the same forward and backward.

// Example 1:

// Input: words = ["abc","car","ada","racecar","cool"]
// Output: "ada"
// Explanation: The first string that is palindromic is "ada".
// Note that "racecar" is also palindromic, but it is not the first.

package main

import "fmt"

func firstPalindrome(words []string) string {

	for _, v := range words {
		v1 := v
		if palindrom(v1) {
			return v1
		}
	}
	return ""
}

func palindrom(v1 string) bool {

	str := ""
	for _, v := range v1 {
		str = string(v) + str
	}
	if v1 == str {
		return true
	} else {
		return false
	}

}

func main() {
	a := []string{"abc", "car", "ada", "racecar", "cool"}
	fmt.Println(firstPalindrome(a))
}
