package main

import "fmt"

type contactInfo struct {
	email string
	address string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	//bruno := person{"Bruno", "Ramalho"}
	//bruno := person{firstName: "Bruno", lastName: "Ramalho"}
	//
	//fmt.Println(bruno)
	//fmt.Println(bruno.firstName + " " + bruno.lastName)
	//fmt.Printf("%+v", bruno)

	//var bruno person
	//bruno.firstName = "Bruno"
	//bruno.lastName = "Ramalho"

	bruno := person{
		firstName: "Bruno",
		lastName: "Ramalho",
		contactInfo: contactInfo{
			email: "bruno@gmail.com",
			address: "Lisbon",
			zipCode: 123,
		},
	}

	//brunoPointer := &bruno
	//brunoPointer.updateName("B")

	bruno.updateName("B")
	bruno.print()
}
