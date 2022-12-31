package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//	type person struct {
//		firstName string
//		lastName  string
//		contact   contactInfo
//	}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact.email = "me@alex.com"
	// alex.contact.zipCode = 12345

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.email = "me@alex.com"
	alex.zipCode = 12345

	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contact: contactInfo{
	// 		email:   "me@jim.com",
	// 		zipCode: 12345,
	// 	},
	// }

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "me@jim.com",
			zipCode: 12345,
		},
	}

	fmt.Printf("%+v\n", alex)
	fmt.Printf("%+v\n", jim)
	jim.print()
	jim.updateName("Jimmy")
	jim.print()

	fmt.Println("Using a pointer receiver")
	jimPointer := &jim
	jimPointer.updateNameUsingPointer("Jimmy pointer")
	jimPointer.print()
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateNameUsingPointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
