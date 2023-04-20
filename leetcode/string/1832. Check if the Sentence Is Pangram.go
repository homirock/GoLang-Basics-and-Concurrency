// A pangram is a sentence where every letter of the English alphabet appears at least once.

// Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.

// Example 1:

// Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
// Output: true
// Explanation: sentence contains at least one of every letter of the English alphabet.

package main

import (
	"fmt"
	"strings"
)

func checkIfPangram(sentence string) bool {
	alp := "abcdefghijklmnopqrstuvwxyz"

	for _, v := range alp {
		if strings.Contains(sentence, string(v)) {
			continue
		} else {
			return false
		}

	}
	return true

}

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
}
