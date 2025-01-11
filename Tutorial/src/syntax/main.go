package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	var s string

	fmt.Println(i, j, k, c, python, java, ToBe, MaxInt, z)
	fmt.Printf("Type: %T Value: %v\n", s, s)
}
