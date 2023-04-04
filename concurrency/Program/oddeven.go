package main

import (
	"fmt"
)

func oddFunc(odd chan int){
	for value:=range odd{ 
		
		fmt.Println("odd:",value)
	}
	
}

func evenFunc(even chan int){
	
	for value:=range even{ 
		
		fmt.Println("even:",value)
	}
		
}

func main(){
	odd := make(chan int)
	even := make(chan int)

	go oddFunc(odd)
	go evenFunc(even)

	for i := 1; i <= 10; i++ {
		if i%2==0{
			even <- i
		}else{
			odd <- i
		}
		
	
}
}