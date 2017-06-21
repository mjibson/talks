package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64 = 2100.825
	fmt.Printf("x: %g\n", x)
	fmt.Printf("r: %g\n", round100(x))
}

func round100(x float64) float64 {
	a, rem := math.Modf(x * 100)
	if rem >= 0.5 {
		a++
	}
	return a / 100
}

// OMIT
