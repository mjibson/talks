package main

import (
	"fmt"

	"github.com/cockroachdb/apd"
)

func main() {
	d := apd.New(27, 0)
	three := apd.New(3, 0)
	c := apd.BaseContext.WithPrecision(5)
	for {
		res, _ := c.Quo(d, d, three)
		fmt.Printf("%7s, inexact: %5v\n", d, res.Inexact())
		if res.Inexact() {
			return
		}
	}
}
