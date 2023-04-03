package main
import "fmt"

func main() {
	ch1:= make(chan int)

	for i := 0; i < 10; i++ {
		ch1<-i
		fmt.Println(<-ch1)
	}
}