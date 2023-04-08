package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	go func(ch1 chan int) {
		for i, j := 0, 1; ; i, j = i+j, i {
			ch1 <- i

		}
	}(ch1)

	for i := 0; i <= 9; i++ {
		fmt.Println(<-ch1)
	}
	close(ch1)

}
