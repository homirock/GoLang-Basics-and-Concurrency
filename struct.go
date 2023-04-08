package main

import "fmt"
//https://medium.com/rungo/structures-in-go-76377cc106a2
func main() {
	type Person struct {
		name string
		id   int
		city string
	}
	//initialize struct with empty value and assigning each value (method1)
	var per1 Person
	per1.name = "pars"
	per1.id = 3
	per1.city = "mp"
	fmt.Println(per1)
	//initialize a stuct with value(method2)
	a1 := Person{name: "homi", id: 2, city: "kjr"}
	//getting the value using dot operator
	fmt.Println("initializing struct:",a1.name, a1.city, a1.id)
	//pointer to the struct
	a2 := &Person{name: "homi", id: 2, city: "kjr"}
	//we havee to use bracket .so so can differntiate betwwne *a2.name and (*a2).name
	fmt.Println("struct with *:",(*a2).name, (*a2).city, (*a2).id)
	//instead of *a2 we can use a2.name directly...go will understand this
	fmt.Println("struct without *:",a2.name, a2.city, a2.id)

	//Anonymous struct
	Person1:= struct {
		name, city string
		id         int
	}{
		name: "Anonymous",
		city: "Anonymous1",
		id:12,
	}

	fmt.Println("Anonymous struct",Person1)

	//Anonymous fields
// You can define a struct type without declaring any field names. You have to just define the field data types and Go will use the data type declarations (keywords) as the field names.

type Person2 struct{
	string
	int
	//bool
}

a3:= &Person2{"partha",1}
fmt.Println("Anonymous field:",a3.string)
//but we cant assign two same type field here
}

// mix of fields 
// type Employee struct {
// 	firstName, lastName string
// 	salary              int
// 	bool                         // anonymous field
// }

//Export struct

//type Employee struct {
	// 	FirstName, LastName string
	// 	salary              int
	// 	bool                         // anonymous field
	// }

	//only FirstName, LastName can be exported outside main function
