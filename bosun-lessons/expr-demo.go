package main

import (
	"fmt"
	"time"

	"bosun.org/cmd/bosun/expr"
)

func main() {
	e, _ := expr.New("7 - 2 < 2 * 3")
	res, _, _ := e.Execute(nil, nil, time.Now(), 0, true, nil, nil, nil)
	r := res.Results[0]
	fmt.Println("value:", r.Value)
	for i, c := range r.Computations {
		fmt.Printf("computation %v: %v = %v\n", i, c.Text, c.Value)
	}
}
