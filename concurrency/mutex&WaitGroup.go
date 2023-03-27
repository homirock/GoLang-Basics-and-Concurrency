package main
import(
	"fmt"
	//"time"
	"sync"
)

// func runner1(wg *sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Println("runner1 stated")
// }

// func runner2(wg *sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Println("runner2 stated")
// }

// func execute(){
// 	//var wg sync.WaitGroup
// 	wg:=new(sync.WaitGroup) //wait group explicitely pass iinto the function by pointer to the sync.WaitGroup struct
// 	wg.Add(2)
// 	go runner1(wg)
// 	go runner2(wg)
// 	wg.Wait()
// }
var m  = 0
func main(){	
	//execute()	
	//time.Sleep(time.Second)
	mu:= new(sync.Mutex)
	wg:= new(sync.WaitGroup)
	wg.Add(3)
	for i:=1;i<4;i++{
	//fmt.Println(i)
	go func(wg *sync.WaitGroup, mu *sync.Mutex){
		mu.Lock()	
		m=m+1		
		mu.Unlock()
		wg.Done()
	}(wg,mu)
	}	
	wg.Wait()
	fmt.Println(m)
}