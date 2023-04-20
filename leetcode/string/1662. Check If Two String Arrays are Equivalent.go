// Given two string arrays word1 and word2, return true if the two arrays represent the same string, and false otherwise.

// A string is represented by an array if the array elements concatenated in order forms the string.

// Example 1:

// Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
// Output: true
// Explanation:
// word1 represents string "ab" + "c" -> "abc"
// word2 represents string "a" + "bc" -> "abc"
// The strings are the same, so return true.
package main

import "fmt"


func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    
	str1:=""
	str2:=""
		for _,v:=range word1{
			str1=str1+v
		}
		for _,v:=range word2{
			str2=str2+v
		}
		if str1==str2{
			return true
		}else{
			return false
		}
	}

	func main(){
		word1:=[]string{"abc", "d"}
		word2:=[]string{"ab", "cd"}
		fmt.Println(arrayStringsAreEqual(word1, word2))
	}