package main

import "fmt"

func main()  {
	//var colors map[string]string
	//colors := make(map[string]string)
	//colors["black"] = "#000"
	//delete(colors, "black")
	//fmt.Println(colors)

	colors := map[string]string{
		"red" : "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	printColors(colors)
}

func printColors(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Color: ", color, " Hex: ", hex)
	}
}
