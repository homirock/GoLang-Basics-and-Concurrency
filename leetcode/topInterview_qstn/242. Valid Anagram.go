// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true
// Example 2:

// Input: s = "rat", t = "car"
// Output: false

package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	s1 := strings.Split(s, "")
	s2 := strings.Split(t, "")
	sort.Strings(s1)
	s3 := strings.Join(s1, "")
	sort.Strings(s2)
	s4 := strings.Join(s2, "")
	if len(s3) == len(s4) {
		if s3 == s4 {
			return true
		}
	}
	return false
}

func isAnagram2(s string, t string) bool {
    if len(s)!=len(t){
        return false
    }
	m:=make(map[rune]int)
for _,v:=range s{
	if _,ok:=m[v];ok{
		m[v]++
	}else{
        m[v]=1
    }
}
fmt.Println(m)
for _,v1:=range t{
	if _,ok:=m[v1];ok{
		m[v1]--
	}
}
for _,v2:=range s{
    if m[v2]!=0{
        return false
    }
}
return true	
}

func main() {
	fmt.Println(isAnagram("", ""))
}
