package main

import (
	"fmt"
)

func oddFunc(odd chan int){
	
	fmt.Println("odd:",<-odd)
	
}

func evenFunc(even chan int){
	
		fmt.Println("even:",<-even)
		
}

func main(){
	odd := make(chan int)
	even := make(chan int)

	go oddFunc(odd)
	go evenFunc(even)

	for i := 1; i <= 10; i++ {
		if i%2==0{
			odd <- i
		}else{
			even <- i
		}
	
}
}