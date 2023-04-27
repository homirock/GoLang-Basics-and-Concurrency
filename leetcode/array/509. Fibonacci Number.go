// The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.
// Given n, calculate F(n).

// Example 1:

// Input: n = 2
// Output: 1
// Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

package main

import "fmt"

func fib(n int) int {
	arr := []int{0,1}
	a := 0
	b := 1
	//l := len(arr)
	for i:=0;i<n-1;i++{
		num := a + b
		a = b
		b = num
		fmt.Println(num)
		arr = append(arr, num)
        

	}
    

	return arr[n-1] + arr[n-2]

}
func main() {
	fib(2)
}
