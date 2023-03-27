package main
import (
	"fmt"
	"time"
)


func channel1( ch1 chan string ){
	time.Sleep(6*time.Second)
	ch1 <- "welcome to ch1"
}

func channel2( ch2 chan string ){
time.Sleep(5*time.Second)
	ch2 <- "welcome to ch2"
}

func main(){
	R1:= make(chan string)
	R2:= make(chan string)
	go channel1(R1)
	go channel2(R2)
	select{
	case op1:= <- R1:
	fmt.Println(op1)

	case op2:= <- R2:
	fmt.Println(op2)
	}

}