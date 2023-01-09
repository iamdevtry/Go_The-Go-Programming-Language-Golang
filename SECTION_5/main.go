package main

import "fmt"

func printColors(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func main() {
	// var colors map[string]string
	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// 	"blue":  "#0000ff",
	// }

	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	colors["red"] = "#ff0000"

	delete(colors, "white")

	// fmt.Println(colors)
	printColors(colors)
}
