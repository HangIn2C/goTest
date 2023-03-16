package main

import "fmt"
// 클로저란 함수를 일급 객체로 취급하는 함수형 프로그래밍 언어에서 사용되는 특성이다.

// Go 함수들은 클로저일 수도 있다.
// 클로저는 함수의 외부로부터 오는 변수를 참조하는 함수 값이다.
// 함수는 참조된 변수에 접근하여 할당될 수 있다.
// 이러한 의미에서 함수는 변수에 "bound" 된다.

// 그 예로, adder 함수는 클로저를 반환한다.
// 각 클로저는 그 자체의 sum 변수에 bound 되어 있다.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
