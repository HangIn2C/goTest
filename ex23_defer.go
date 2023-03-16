package main

import (
	"fmt"
)

func test() {
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}
}

func main() {
	// defer문은 자신을 둘러싼 함수가 종료할 때까지 어떠한 함수의 실행을 연기
	defer fmt.Println("world")

	fmt.Println("hello")

	test()
}
