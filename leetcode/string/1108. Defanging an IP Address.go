// Given a valid (IPv4) IP address, return a defanged version of that IP address.

// A defanged IP address replaces every period "." with "[.]".
// Example 1:

// Input: address = "1.1.1.1"
// Output: "1[.]1[.]1[.]1"
// Example 2:

// Input: address = "255.100.50.0"
// Output: "255[.]100[.]50[.]0"

package main
import (
	"fmt"
	s"strings"
)

func main() {
str:="1.2.3"

fmt.Println(s.ReplaceAll(str,".","[.]"))
split:=s.Split(str,".")
fmt.Println(split)
fmt.Println(s.Join(split,"[.]"))
}