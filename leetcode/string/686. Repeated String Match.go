// Given two strings a and b, return the minimum number of times you should repeat string a so that string b is a substring of it. If it is impossible for b​​​​​​ to be a substring of a after repeating it, return -1.

// Notice: string "abc" repeated 0 times is "", repeated 1 time is "abc" and repeated 2 times is "abcabc".

// Example 1:

// Input: a = "abcd", b = "cdabcdab"
// Output: 3
// Explanation: We return 3 because by repeating a three times "abcdabcdabcd", b is a substring of it.
package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	str := ""
	count := 0
	for {
		count = count + 1
		str = str + a
		if strings.Contains(str, b) {
			return count

		} else if len(str) > (2 * len(b)) {
			return -1
		}
	}
}
func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
}
