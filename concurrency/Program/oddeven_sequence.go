package main

import (
	"fmt"
)

func oddEven(ch1 chan int,quit chan bool) {
			for v := range ch1 {
				if v%2 == 0 {
					fmt.Println("even:", v)
				} else {
					fmt.Println("odd:", v)
				}
			}
			quit<-true
}

func main() {
	ch1 := make(chan int)
	quit := make(chan bool)

	arr := []int{0,1, 2, 3, 4, 5, 6, 7, 8}

	l := len(arr)

	go oddEven(ch1,quit)
	go func() {
		for i := 0; i <= l-1; i++ {
			ch1 <- arr[i]
		}
		quit<-false
	}()
	<-quit

}
