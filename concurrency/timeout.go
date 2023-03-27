// For our example, suppose we’re executing an external call that returns its result on a channel c1 after 2s. Note that the channel is buffered, so the send in the goroutine is nonblocking. This is a common pattern to prevent goroutine leaks in case the channel is never read.
package main
import (
	"fmt"
	"time"
)

func func1(op1 chan string){
	time.Sleep(2*time.Second)
	op1<- "result1"
}

func func2(op2 chan string){
	time.Sleep(1*time.Second)
	op2<- "result2"
}


func main(){
	op1:=make(chan string,1) //non-blocking channel

	go func1(op1)

	select{
	case res:=<-op1:
	fmt.Println(res)
	case <-time.After(1*time.Second):
	fmt.Println("timeout1")
	}
	op2:=make(chan string,1) //non-blocking channel
	go func2(op2)
	select{
	case res1:=<-op2:
	fmt.Println(res1)
	case <-time.After(3*time.Second):
	fmt.Println("timeout2")
	}
}