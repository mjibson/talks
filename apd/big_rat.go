package main

import (
	"fmt"
	"math/big"
)

func main() {
	r := new(big.Rat)
	r.SetString("123e40")
	fmt.Println(r)
}
