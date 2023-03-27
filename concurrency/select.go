package main
import (
	"fmt"
	"time"
)

func main(){
	c1:= make(chan string)
	c2:= make(chan string)

	go func(){
		//time.Sleep(5*time.Second)
		c1 <- "message1"
	}()
	go func(){
		time.Sleep(3*time.Second)
		c1 <- "message2"
	}()
 for i:=0;i<2;i++{
	select{
	case msg1:= <-c1:
	fmt.Println(msg1)
	case msg2:= <-c2:
	fmt.Println(msg2)
	default:
		fmt.Println("not executed")
	}
}

}

//if no values in any of the channel then it will execute the default value.

// ex:- output--not executed
// 			message1