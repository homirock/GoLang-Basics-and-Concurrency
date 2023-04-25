// Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

// Example 1:

// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	stringArray := []string{}

	stringArr1 := strings.Split(s, " ")
	for _, v1 := range stringArr1 {
		rev := reverseString(v1)
		stringArray = append(stringArray, rev)
	}
	finalstr := strings.Join(stringArray, " ")
	return finalstr
}

func reverseString(v1 string) string {
	str1 := ""
	for _, v2 := range v1 {
		str1 = string(v2) + str1
	}
	return str1
}

func main() {
	fmt.Println(reverseWords("Hi Partha Sarathi"))
}
