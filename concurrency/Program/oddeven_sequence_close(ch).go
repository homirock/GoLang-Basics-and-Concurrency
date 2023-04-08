package main

import (
	"fmt"
)

func oddEven(ch1 chan int) {
	for v := range ch1 {
		if v%2 == 0 {
			fmt.Println("even:", v)
		} else {
			fmt.Println("odd:", v)
		}
	}
}

func main() {
	ch1 := make(chan int)
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	l := len(arr)
	go oddEven(ch1)
	func() {
		for i := 0; i <= l-1; i++ {
			ch1 <- arr[i]
		}
		close(ch1)
	}()
}
