// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

// Example 1:

// Input: nums = [3,0,1]
// Output: 2
// Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the //missing number in the range since it does not appear in nums.

package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {

	val := 0
	l := len(nums)
	sort.Ints(nums)
	if nums[l-1] != l {
		return l
	}
	for k, v := range nums {
		if k != v {
			val = k
			break
		}
	}
	return val

}

func main() {
	arr := []int{1,4,3,0}
	fmt.Println(missingNumber(arr))
}
