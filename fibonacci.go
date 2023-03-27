// dynamic fibonacci series
package main
import(
	"fmt"
	"strings"
	"strconv"
)

func fibonacci (first,second ,num int){
	// n1:=first
	// n2:=second
	s1:= strconv.Itoa(first)
	s2:= strconv.Itoa(second)
	arr := []string{s1,s2}
	//fmt.Printf("%d%d",n1,n2)
	for i:=0;i<=num;i++{
     n3:= first+second
	 first = second
	 second = n3
	 s3:= strconv.Itoa(n3)
	 arr=append(arr,s3)
	//fmt.Printf("%d",n3)
	//fmt.Println(arr)
	}
	fmt.Println(arr)
	x:=strings.Join(arr,",")
	fmt.Printf("%v",[]string{x})
}

func main(){
	fibonacci(2,5,10)
}