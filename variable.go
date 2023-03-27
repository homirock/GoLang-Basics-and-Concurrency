//In Go, variables are explicitly declared and used by the compiler to e.g. 
//check type-correctness of function calls.
//Variables declared without a corresponding initialization are zero-valued. 
//For example, the zero value for an int is 0.

//The := syntax is shorthand for declaring and initializing a variable,
// e.g. for var f string = "apple" in this case.
package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
}