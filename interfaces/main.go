package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
func (englishBot) getGreeting() string {
	return "Hello!"
}

type portugueseBot struct {}
func (portugueseBot) getGreeting() string {
	return "Olá!"
}

func printGreeting(b bot)  {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	pb := portugueseBot{}

	printGreeting(eb)
	printGreeting(pb)
}
