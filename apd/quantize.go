package main

import (
	"fmt"

	"github.com/cockroachdb/apd"
)

func main() {
	input, _, _ := apd.NewFromString("123.45")
	output := new(apd.Decimal)
	c := apd.BaseContext.WithPrecision(10)
	for i := int32(-3); i <= 3; i++ {
		res, _ := c.Quantize(output, input, i)
		fmt.Printf("%2v: %s %s\n", i, output, res)
	}
}
