package main

import (
	"fmt"
	"sync"
)



func sendData(ch chan int, v int, wg *sync.WaitGroup, m *sync.Mutex) {
	
	m.Lock()	
	ch <- v
	value := <-ch
	// fmt.Println(value)
	if value%2 == 0 {
		fmt.Println("even value:",value)
			
	}else{
		fmt.Println("odd value:",value)
	}
	 m.Unlock()
	 wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	arr := []int{1, 2, 3, 4, 5, 6}
	
	wg.Add(6)
	for _, v := range arr {
		ch := make(chan int,1)
		go sendData(ch, v,&wg,&m)
		
	}
	wg.Wait()
	
}
