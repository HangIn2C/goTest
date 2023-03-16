package main

import "fmt"

// 0 메서드를  지칭하는 인터페이스 유형을 _empty_interface라고 합니다.
// interface{}
// 빈 인터페이스는 모든 유형의 값을 가질 수 있습니다.
// (모든 유형은 최소 0개의 메소드를 구현합니다.)

// 빈 인터페이스는 알 수 없는 값을 처리하는데 이용됩니다.
// 예를 들어 fmt.print는 interface{} 타입의 어떠한 인수라도 가질 수 있다.

func main() {
	// 초기에는 빈 인터페이스 이기 때문에 nil
	var i interface{}
	describe(i)

	// i의 값으로 42를 넣으면 빈 인터페이스는 int 타입으로 변경
	i = 42
	describe(i)

	// i의 값으로 "hello"를 넣으면 빈 인터페이스는 string 타입으로 변경
	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
