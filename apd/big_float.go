package main

import (
	"fmt"
	"math/big"
)

func main() {
	const precision = 10
	d1, _, _ := big.ParseFloat("0.1", 10, precision, big.ToNearestEven)
	d2 := new(big.Float).Copy(d1)
	for i := 1; i < 7; i++ {
		d2.Add(d2, d1)
		fmt.Println(i, d2)
	}
}
