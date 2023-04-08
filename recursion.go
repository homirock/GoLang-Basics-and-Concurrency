package main

import "fmt"

//import "fmt"

func factorial(n int) int {
	if n==1{
		return 1
	}
	fmt.Println(n)
	return n * factorial(n-1)
}
func  anonymous() func()int{
	i := 0
    return func() int {
        i++
        return i
    }
}
func main() {
	//var fib func()
	fmt.Println(factorial(6))
	intf:=anonymous()
	fmt.Println(intf())

	

}
