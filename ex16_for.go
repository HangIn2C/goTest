package main

import "fmt"

func main() {
	// Go는 for 반복문 하나의 반복 구조를 가진다.
	// for 초기화 구문, 조건 표현, 사후 구문
	// 위의 형식을 가진다.
	// Go는 다른 언어와 달리 구성 요소를 () 감싸지 않고, { } 괄호가 필수이다.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(i)
	}
	fmt.Println(sum)

	// 초기화 구문과 사후 구문은 필수가 아니다.
	sum2 := 1
	for ; sum < 1000; {
		sum2 += sum
	}
	fmt.Println(sum2)

	// ;을 생략 하는점에서 while이 필요없게 된다.
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	// for문 무한루프
	// for {
	// }
}
