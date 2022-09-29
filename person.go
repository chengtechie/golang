package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func (p *person) print() {
	fmt.Printf("%+v", *p)
}

func printPeople() {
	var human2 person
	human2.firstName = "Kim"
	human2.lastName = "Hor"
	human2.contactInfo.email = "kim@yahoo.com"
	human2Pointer := &human2
	human2Pointer.updateName("Kiwi")
	human2.print()

	human1 := person{
		firstName: "Cheng",
		lastName:  "Lim",
		contactInfo: contactInfo{
			email: "cheng@gmail.com",
		},
	}
	human1.print()
}
