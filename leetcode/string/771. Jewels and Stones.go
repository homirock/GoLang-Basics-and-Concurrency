// You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have. Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.

// Letters are case sensitive, so "a" is considered a different type of stone from "A".

// Example 1:

// Input: jewels = "aA", stones = "aAAbbbb"
// Output: 3
// Example 2:

// Input: jewels = "z", stones = "ZZ"
// Output: 0

package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	count := 0
	for _, v := range jewels {

		ltr := string(v)

		for _, v := range stones {
			ltr1 := string(v)
			if ltr == ltr1 {
				count = count + 1
			}
		}
	}
	return count

}

func main() {
	fmt.Println(numJewelsInStones("aAB","aaaaAAAABBB"))
}
