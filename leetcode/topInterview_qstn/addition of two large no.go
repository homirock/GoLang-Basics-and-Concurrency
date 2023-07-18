package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "123"
	y := "128"
	var l int
	var m string
	var n string
	if x > y {
		l = len(x) - 1
		m = x
		n = y
	} else {
		l = len(y) - 1
		m = y
		n = x
	}
	new := ""
	var carry int
	for i := l; i >= 0; i-- {
		a, _ := strconv.Atoi(string(m[i]))
		b, _ := strconv.Atoi(string(n[i]))
		c := a + b + carry
		digit := c % 10
		if c < 10 {
			new = strconv.Itoa(digit) + new
			carry = 0
		} else {
			new = strconv.Itoa(digit) + new
			carry = c / 10
		}

	}
	res, _ := strconv.Atoi(new)
	fmt.Println(res)
}
