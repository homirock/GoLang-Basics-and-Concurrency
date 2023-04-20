// A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each word consists of lowercase and uppercase English letters.

// A sentence can be shuffled by appending the 1-indexed word position to each word then rearranging the words in the sentence.

// For example, the sentence "This is a sentence" can be shuffled as "sentence4 a3 is2 This1" or "is2 sentence4 This1 a3".
// Given a shuffled sentence s containing no more than 9 words, reconstruct and return the original sentence.

// Example 1:

// Input: s = "is2 sentence4 This1 a3"
// Output: "This is a sentence"
// Explanation: Sort the words in s to their original positions "This1 is2 a3 sentence4", then remove the numbers.

package string

import (
	"fmt"
	"strings"
)

func sortSentence(s string) string {
	arr := strings.Split(s, " ")
	for i := 0; i < len(arr); i++ {
		// fmt.Println("byte value of ", arr[i][len(arr[i])-1], " -- ", arr[i])
		val := arr[i][len(arr[i])-1] - '0'
		// fmt.Println(val, " --- ", arr)
		for j := i + 1; j < len(arr); j++ {
			val1 := arr[j][len(arr[j])-1] - '0'
			if val > val1 {
				val = val1
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	// fmt.Println(arr)
	var rslt string
	for i := 0; i < len(arr); i++ {
		rslt += string(arr[i][:len(arr[i])-1])
		if i != len(arr)-1 {
			rslt += " "
		}
	}
	return rslt
}

func main() {
	fmt.Println(sortSentence("My1 Name2 is3 Partha4"))
}
