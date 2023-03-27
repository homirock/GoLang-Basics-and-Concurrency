//The function that is called with the varying number of arguments is known as variadic function.
// Or in other words, a user is allowed to pass zero or more arguments in the variadic function. 
//fmt.Printf,Println is an example of the variadic function, it required one fixed argument at the starting after 
//that it can accept any number of arguments. 

//Important Points: In the declaration of the variadic function, the type of the last parameter is preceded by an ellipsis,
// i.e, (…). It indicates that the function can be called at any number of parameters of this type. 

//Syntax:  

//function function_name(para1, para2...type)type {// code...}


package main
import(
	"fmt"
)

func variadicfunction(nums ...int){
	count := 0
    for _,v := range nums{
		count++
		fmt.Println(v)
	}
	fmt.Println(count)
}

func main(){
	//fmt.Println("1","2")
	variadicfunction(8,9)
	// If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	variadicfunction(nums...)
}