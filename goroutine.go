package main

import (
"fmt"
"time"
)

func function1(text string){
	for i:=0;i<=5;i++{
		fmt.Println(text)
	}
}
func main(){
	//function1("normal function")
	go function1("go routine")
	time.Sleep(10)
	//function1("normal function")

}