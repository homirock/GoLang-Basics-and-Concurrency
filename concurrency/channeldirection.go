package main
import "fmt"

func pings(ping chan string, message string){
	ping<- message
	}

func pongs(ping chan string ,pong chan string){
	message1,state:= <-ping
	fmt.Println(state)
	pong<-message1
}

func main(){

	ping := make (chan string)
	pong := make (chan string)
	message:="ping"
	go pings(ping,message)
	go pongs(ping,pong)
	fmt.Println(<-pong)
}