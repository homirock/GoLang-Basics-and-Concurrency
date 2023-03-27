//https://medium.com/bootdotdev/golang-mutexes-what-is-rwmutex-for-5360ab082626

package main
import (
	"fmt"
	"sync"
)
// By using a sync.RWMutex, our program becomes more efficient. We can have as many readLoops as we want to access our data, but we can assure that the writers have exclusive access.

 func writeMap(m map[int]int,mu *sync.RWMutex){
	for{
	for i:=0;i<10;i++{
		mu.Lock()
		m[i]=i
		mu.Unlock()
	}
	fmt.Println(m)
}
}

func readMap(m map[int]int,mu *sync.RWMutex){
	for{
		mu.RLock()
	for k,v:= range m{	
		fmt.Println(k,"-",v)	
	}
	mu.RUnlock()
	}
}

func main(){
	fmt.Println("started")
	channel:= make(chan struct{})
	var mu sync.RWMutex
	fmt.Println("in progress")
	m:= map[int]int{}
	go writeMap(m,&mu)
	go readMap(m,&mu)
	fmt.Println("in progress run")
	<- channel
	fmt.Println("terminated")
}

