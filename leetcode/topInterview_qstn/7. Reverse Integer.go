// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:

// Input: x = 123
// Output: 321
// Example 2:

// Input: x = -123
// Output: -321
// Example 3:

// Input: x = 120
// Output: 21
package main

import "fmt"

func main() {
	fmt.Println(reverse(123))
}
func reverse(x int) int {
	num := x
	flag := false
	if num < 0 {
		num = num * (-1)
		flag = true

	}
	result := 0
	for num > 0 {
		remainder := num % 10
		result = result*10 + remainder
		num = num / 10
	}
	if !flag {
		return result
	}
	return result * (-1)

}
