package main

import "fmt"

const PI = 3.14

func main() {
	// 상수는 := 를 통해 선언될 수 없다
	const WORLD = "world"
	fmt.Println("Hello", WORLD)
	fmt.Println("Happy", PI, "Day")

	const TRUTH = true
	fmt.Println("Go rules?", TRUTH)
}
