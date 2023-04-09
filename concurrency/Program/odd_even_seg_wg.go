package main

import (
	"fmt"
	"sync"
)

func oddEven(ch1 chan int, wg *sync.WaitGroup) {

	for value := range ch1 {
		if value%2 == 0 {
			fmt.Println("even:", value)
		} else {
			fmt.Println("odd:", value)
		}

	}

	wg.Done()
	//quit <- true

}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	ch1 := make(chan int)
	var wg sync.WaitGroup
	//quit := make(chan bool)
	wg.Add(2)
	go oddEven(ch1, &wg)
	go func(wg *sync.WaitGroup) {
		for _, v := range arr {
			ch1 <- v
		}
		close(ch1)
		//time.Sleep(3 * time.Second)
		wg.Done()
		//quit <- true
	}(&wg)
	wg.Wait()
	//<-quit
}
