package main

import (
	"fmt"
	"math"
)

func add(ns ...float64) float64 {
	var total float64

	for _, n := range ns {
		total += n
	}

	return total
}

func sub(x float64, ns ...float64) float64 {
	for _, n := range ns {
		x -= n
	}

	return x
}

func mul(x float64, ns ...float64) float64 {
	for _, n := range ns {
		x *= n
	}

	return x
}

func div(x float64, ns ...float64) float64 {
	if x == 0 {
		return x
	}

	for _, n := range ns {
		if n == 0 {
			fmt.Println("Numbers is not able to divide by 0")
			continue
		}

		x /= n
	}

	return x
}

func pow(x float64, p float64) float64 {
	return math.Pow(x, p)
}
