// Write an algorithm to determine if a number n is happy.

// A happy number is a number defined by the following process:

// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.

// Example 1:

// Input: n = 19
// Output: true
// Explanation:
// 12 + 92 = 82
// 82 + 22 = 68
// 62 + 82 = 100
// 12 + 02 + 02 = 1
package main

import "fmt"

func main() {
	fmt.Println(isHappy(123))
}
func isHappy(n int) bool {
	count := 0
	for {
		count++
		result := 0
		for n > 0 {
			remainder := n % 10
			result = result + remainder*remainder
			n = n / 10
		}
		if result == 1 {
			return true
		} else if result == n {
			return false
		}
		n = result
		if count == 20 {
			return false
		}
	}

}
