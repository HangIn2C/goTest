package main

import "fmt"

func main() {
	// 슬라이스에 새로운 요소를 추가하는 것이 일반적
	// Go는 내장된 append 함수를 제공한다.
	// append의 첫번째 파라미터 s는 슬라이스의 타입 T이다.
	// 그리고 나마지 T 값들은 슬라이스에 추가할 값들이다.
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
