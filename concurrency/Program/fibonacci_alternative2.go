package main

import (
	"fmt"
)

func fibonacci(ch1 chan int, quit chan bool) {
	count := 0
	for value := range ch1 {
		count = count + 1
		fmt.Println(value)
		if count == 6 {
			quit <- true
		}
	}

}

func main() {
	ch1 := make(chan int)
	quit := make(chan bool)
	go fibonacci(ch1, quit)

	go func(ch1 chan int) {
		//count := 0
		for i, j := 0, 1; ; i, j = i+j, i {
			ch1 <- i
		}

	}(ch1)
	<-quit
}
