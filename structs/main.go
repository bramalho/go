package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	//bruno := person{"Bruno", "Ramalho"}
	//bruno := person{firstName: "Bruno", lastName: "Ramalho"}
	//
	//fmt.Println(bruno)
	//fmt.Println(bruno.firstName + " " + bruno.lastName)

	var bruno person
	bruno.firstName = "Bruno"
	bruno.lastName = "Ramalho"

	fmt.Printf("%+v", bruno)
}
