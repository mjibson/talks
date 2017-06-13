package main

import (
	"fmt"

	"github.com/cockroachdb/apd"
)

func main() {
	// Create a context that will overflow at 1e3.
	c := apd.Context{
		MaxExponent: 2,
		Traps:       apd.Overflow,
	}
	one := apd.New(1, 0)
	d := apd.New(997, 0)
	for {
		_, err := c.Add(d, d, one)
		fmt.Printf("%8s, err: %v\n", d, err)
		if err != nil {
			return
		}
	}
}
