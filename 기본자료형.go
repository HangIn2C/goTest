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

var (
	zero int    // 초기값 0
	lamp bool   // 초기값 false
	str  string // 초기값 "" 빈문자열
)

func main() {
	fmt.Printf("Type : %T Value : %v\n", ToBe, ToBe)
	fmt.Printf("Type : %T Value : %v\n", MaxInt, MaxInt)
	fmt.Printf("Type : %T Value : %v\n", z, z)

	fmt.Println(zero, lamp, str)
}
