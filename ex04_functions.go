package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func mul(x float32, y float32) float32 {
	return x * y
}

func div(x float32, y float32) float32 {
	return x / y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))
	fmt.Println(mul(42, 13))
	fmt.Println(div(42, 13))
}
