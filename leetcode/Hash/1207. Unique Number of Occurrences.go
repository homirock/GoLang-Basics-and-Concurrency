// Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.

// Example 1:

// Input: arr = [1,2,2,1,1,3]
// Output: true
// Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
// Example 2:

// Input: arr = [1,2]
// Output: false
// Example 3:

// Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
// Output: true
package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 3, 4}))
}
func uniqueOccurrences(arr []int) bool {
	var m = make(map[int]int)
	var n = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := m[arr[i]]; ok {
			m[arr[i]]++
		} else {
			m[arr[i]] = 1
		}
	}
	for _, v := range m {
		if _, ok := n[v]; ok {
			n[v]++
			if n[v] >= 2 {
				return false
			}
		} else {
			n[v] = 1
		}
	}
	return true

}
