package main 
import "fmt"

func main(){
	channel:= make (chan string,2)
	channel <-"message1"
	channel <-"message2"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}