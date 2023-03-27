//https://medium.com/bootdotdev/golang-mutexes-what-is-rwmutex-for-5360ab082626

package main
import (
	"fmt"
	"sync"
)

// In the code above we add a sync.Mutex{} and name it mux. In the write loop, we Lock() the mutex before writing, and Unlock() it when we are done. This ensures that no other threads can Lock() the mutex while we have it locked — those threads will block and wait until we Unlock().

// In the reading loop we Lock() before iterating over the map, and likewise Unlock() when we are done.
 func writeMap(m map[int]int,mu *sync.Mutex){
	for{
	for i:=0;i<10;i++{
		mu.Lock()
		m[i]=i
		mu.Unlock()
	}
	fmt.Println(m)
}
}

func readMap(m map[int]int,mu *sync.Mutex){
	for{
		mu.Lock()
	for k,v:= range m{	
		fmt.Println(k,"-",v)	
	}
	mu.Unlock()
	}
}

func main(){
	fmt.Println("started")
	channel:= make(chan struct{})
	var mu sync.Mutex
	fmt.Println("in progress")
	m:= map[int]int{}
	go writeMap(m,&mu)
	go readMap(m,&mu)
	fmt.Println("in progress run")
	<- channel
	fmt.Println("terminated")
}

// By using a sync.RWMutex, our program becomes more efficient. We can have as many readLoops as we want to access our data, but we can assure that the writers have exclusive access.