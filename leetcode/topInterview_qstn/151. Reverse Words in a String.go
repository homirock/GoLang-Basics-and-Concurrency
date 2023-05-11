// Given an input string s, reverse the order of the words.

// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

// Return a string of the words in reverse order concatenated by a single space.

// Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

// Example 1:

// Input: s = "the sky is blue"
// Output: "blue is sky the"
// Example 2:

// Input: s = "  hello world  "
// Output: "world hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.
// Example 3:

// Input: s = "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	trimmed := strings.TrimSpace(s)
	str := strings.Join(strings.Fields(trimmed), " ")
	arr := strings.Split(str, " ")

	l := len(arr)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {

		arr[i], arr[j] = arr[j], arr[i]

	}
	return strings.Join(arr, " ")
}

func main() {
	fmt.Println(reverseWords("a good   example"))
}
