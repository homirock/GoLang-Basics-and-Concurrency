// Given an array of integers nums, return the number of good pairs.

// A pair (i, j) is called good if nums[i] == nums[j] and i < j.

// Example 1:

// Input: nums = [1,2,3,1,1,3]
// Output: 4
// Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.

package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	count := 0
	l := len(nums) - 1
	for i := 0; i <= l-1; i++ {
		for j := i + 1; j <= l; j++ {
			if nums[i] == nums[j] {
				count = count + 1
			}
		}

	}
	return count
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 4}))
}
