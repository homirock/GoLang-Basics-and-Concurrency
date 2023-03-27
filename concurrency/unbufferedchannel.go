package main
import( 
	"fmt"
)

func channelfun(message chan string,message1 string){
	fmt.Println("started")
	message <- message1
	fmt.Println("finished")
	//wont work bcz when we send message to the ch, the receiver should be ready.
	// receive:= <-message
	
	// fmt.Println(receive)

}

func main(){
message:= make(chan string)
	message1:="ping"
	go channelfun(message,message1)
	receive:= <-message
	fmt.Println(receive)
}