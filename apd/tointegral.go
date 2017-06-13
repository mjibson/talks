package main

import (
	"fmt"

	"github.com/cockroachdb/apd"
)

func main() {
	inputs := []string{
		"123.4",
		"123.0",
		"123",
		"12E1",
		"120E-1",
		"120E-2",
	}
	for _, input := range inputs {
		d, _, _ := apd.NewFromString(input)
		res, _ := apd.BaseContext.ToIntegralX(d, d)
		fmt.Printf("% 6s -> % 3s: %s\n", input, d, res)
	}
}
