// Given an alphanumeric string s, return the second largest numerical digit that appears in s, or -1 if it does not exist.




// An alphanumeric string is a string consisting of lowercase English letters and digits.




// Example 1:




// Input: s = "dfa12321afd"

// Output: 2

// Explanation: The digits that appear in s are [1, 2, 3]. The second largest digit is 2.

// Example 2:




// Input: s = "abc1111"

// Output: -1

// Explanation: The digits that appear in s are [1]. There is no second largest digit.

package main




import (

"fmt"

"strconv"

)




func secondHighest(s string) int {
 largest := byte('0')
 secLargest := byte('0')
 count := 0
str := []byte{}
for i := 0; i < len(s); i++ {

if s[i] >= 48 && s[i] <= 57 {

          str = append(str, s[i])

            if s[i] > largest {

                largest = s[i]

            }

        }

    }

    for j := 0; j < len(str); j++ {

        if str[j] == str[0] {

            count++

        }

        if str[j] > secLargest && str[j] != largest {

            secLargest = str[j]

        }

    }

    if len(str) == count {

        return -1

    }

    strval := string(secLargest)

    v, _ := strconv.Atoi(strval)

    return v

}




func main() {

    fmt.Println(secondHighest("asfg6745"))

}