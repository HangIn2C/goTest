package main

import "fmt"

func main() {
	// var a[10]int
	// 위 표현은 변수 a를 10개의 정수들의 배열로 선언한 것이다.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
