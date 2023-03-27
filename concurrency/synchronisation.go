package main
import "fmt"

func worker(done chan bool){

fmt.Println("started")
fmt.Println("end")
done <- true
}

func main(){
	done:= make(chan bool,1)

	go worker(done)
// if we wont receiver here then the program might end before execution of go routine.
	<-done
}