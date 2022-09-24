package main

import "fmt"

func main() {
	intSlices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, slice := range intSlices {
		if slice%2 == 0 {
			fmt.Println(slice, "is even")
			continue
		}

		fmt.Println(slice, "is odd")
	}
}
