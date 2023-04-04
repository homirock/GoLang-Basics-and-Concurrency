package main

import (
	"fmt"
	"sync"
)



func sendData(ch chan int, v int, wg *sync.WaitGroup, m *sync.Mutex) {
	
		
	ch <- v
	wg.Done()
}

func main() {

	count:=0
	var wg sync.WaitGroup
	var m sync.Mutex
	arr := []int{1, 2, 3, 4, 5, 6}
	ch := make(chan int)
	wg.Add(6)
	for _, v := range arr {	

		go sendData(ch, v,&wg,&m)

		
	}

	for{
if count==len(arr){
	break
}
	value := <-ch
	count=count+1
	fmt.Println(value)
	if value%2 == 0 {
		fmt.Println("even value:",value)
			
	}else{
		fmt.Println("odd value:",value)
	}
}
	 
	wg.Wait()
	
}
