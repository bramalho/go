package main

import "fmt"

type bot interface {
	getGreeting() string
}
func printGreeting(b bot)  {
	fmt.Println(b.getGreeting())
}

type englishBot struct {}
func (englishBot) getGreeting() string {
	return "Hello!"
}

type portugueseBot struct {}
func (portugueseBot) getGreeting() string {
	return "Olá!"
}

func main() {
	eb := englishBot{}
	pb := portugueseBot{}

	printGreeting(eb)
	printGreeting(pb)
}
