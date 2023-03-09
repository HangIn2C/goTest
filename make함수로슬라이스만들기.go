package main

import "fmt"

func main() {
	// 슬라이스는 내장된 make 함수로 생성할 수 있다.
	// 이것은 동적 크기의 배열을 생성하는 방법입니다.
	// make 함수는 0으로 이루어진 배열을 할당합니다.
	// 그리고 해당 배열을 참조하는 슬라이스를 반환 합니다.

	a := make([]int, 5) // len(a) = 5
	printSlice("a", a)

	// b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5
	// b[:cap(b)] // len(b)=5, cap(b)=5
	// b = b[1:] // len(b)=4, cap(b)=4
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len = %d cap = %d %v\n", s, len(x), cap(x), x)
}
