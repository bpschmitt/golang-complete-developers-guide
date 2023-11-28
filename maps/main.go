package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := make(map[int]string)

	// colors[10] = "#FFFFFF"

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
	}

	// delete(colors, 10)

	printMap(colors)

}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("%s has a hex value of %s\n", color, hex)
	}
}
