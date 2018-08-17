package main

import "fmt"

func main()  {
	//var colors map[string]string
	//colors := make(map[string]string)

	colors := map[string]string{
		"red" : "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	colors["black"] = "#000"

	delete(colors, "black")

	fmt.Println(colors)
}
