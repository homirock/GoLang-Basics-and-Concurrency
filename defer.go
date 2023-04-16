package main

import "fmt"

func main() {

	defer fmt.Println("started")
	defer fmt.Println("in progress")
	defer fmt.Println("finished")
}