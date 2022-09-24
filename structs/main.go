package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var john person
	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(alex.firstName, alex.lastName)

	john.firstName = "John"
	john.lastName = "Smith"

	fmt.Printf("%+v", john)
}
