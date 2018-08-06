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
	contact contactInfo
}

func main() {
	//bruno := person{"Bruno", "Ramalho"}
	//bruno := person{firstName: "Bruno", lastName: "Ramalho"}
	//
	//fmt.Println(bruno)
	//fmt.Println(bruno.firstName + " " + bruno.lastName)

	//var bruno person
	//bruno.firstName = "Bruno"
	//bruno.lastName = "Ramalho"

	bruno := person{
		firstName: "Bruno",
		lastName: "Ramalho",
		contact: contactInfo{
			email: "bruno@gmail.com",
			address: "Lisbon",
			zipCode: 123,
		},
	}

	fmt.Printf("%+v", bruno)
}
