package main

import (
	"fmt"
	"sync"
)

func main() {
	num := 21
	chO := make(chan int)
	chE := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go odd( num,chO, chE, &wg)
	go even(num, chO, chE, &wg)

	<-chE
	wg.Wait()

}

func odd( num int,chO chan int, chE chan int, wg *sync.WaitGroup) {
	onum := 1
	for onum <= num {

		chO <- onum
		fmt.Println("odd num ", onum)
		onum += 2
		<-chE
	}
	close(chO)
	wg.Done()
}

func even(num int, chO chan int, chE chan int, wg *sync.WaitGroup) {
	enum := 0
	for enum <= num {
		chE <- enum
		fmt.Println("even num ", enum)
		enum += 2
		<-chO
	}
	close(chE)
	wg.Done()
}
