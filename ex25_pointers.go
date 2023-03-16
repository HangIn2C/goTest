package main

import "fmt"

func main() {
	// 포인터는 값의 메모리 주소를 가지고 있다.
	// *T 타입은 T 값을 가리키는 포인터이고 초기값은 nil(=null)
	
	// var p *int

	i, j := 42, 2701

	// & 연산자는 이것의 피연산자에 대한 포인터를 생성합니다.
	// i := 42
	// p = &i
	p := &i // 42 // p에 포인터 생성
	fmt.Println("pointer i : ", *p)

	// * 연산자는 포인터가 가리키는 주소의 값을 나타냅니다.
	// fmt.Println(*p) // 포인터 p를 통해 i 읽기
	// *p = 21
	// 역참조 또는 간접 참조라 한다.
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
