package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode uint
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// var john person
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// fmt.Println(alex.firstName, alex.lastName)

	// john.firstName = "John"
	// john.lastName = "Smith"

	// fmt.Printf("%+v", john)

	var jim person = person{
		firstName: "Jim",
		lastName:  "J.",
		contact:   contactInfo{email: "jim@gmail.com", zipCode: 1234},
	}

	jim.setFirstName("First")
	jim.print()
}

func (p *person) setFirstName(firstName string) {
	p.firstName = firstName
}

func (p person) setLastName(lastName string) {
	p.lastName = lastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
