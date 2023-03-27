package main
import (
	"fmt"
)

func main(){
	arr := []int{1,2,3,4,5}

	for k,v:= range arr {
		fmt.Printf("%v:%v\n",k,v)
	}
}