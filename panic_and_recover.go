package main

import "fmt"

func callName(name *string, age *string) {
	defer fmt.Println("starting of callName")
	if name == nil {
		panic("Name is not mentioned")
	}
	defer fmt.Println("In between callName")
	if age == nil {
		panic("Age is not mentioned")
	}
	defer fmt.Println("Ending of callName")
}

func main() {
	defer fmt.Println("starting of Main function.")
	name := "partha"

	callName(&name, nil)
}
