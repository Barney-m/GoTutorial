package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	// fmt.Println(colors["white"])

	printMap(colors)
}

func printMap(colors map[string]string) {
	for key, value := range colors {
		fmt.Println(key, value)
	}
}
