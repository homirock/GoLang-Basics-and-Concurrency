// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s1 := strings.ToLower(s)
	str := ""
	str1 := ""
	for i := 0; i < len(s1); i++ {
		if s1[i] >= 97 && s1[i] <= 122 || s1[i] >= 48 && s1[i] <= 57 {
			str1 = string(s1[i]) + str1 //reverse string
			str = str + string(s1[i])   //actual string
		}
	}

	if str1 == str {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
