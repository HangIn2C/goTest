package main

import "fmt"

func main() {
	primes := [6]int{2,3,5,7,11,13}

	// 배열은 고정된 크기를 가진다.
	// 반면 슬라이스는 배열의 요소들을 동적인 크기로, 유연하게 볼 수 있다.
	// 슬라이스는 콜론으로 구분된 두 개의 인덱스(하한과 상한)를 지정하여 형성
	// a[low : high]
	// a[1 : 4]
	// 위의 표현은 a의 첫번째부터 세번째 요소를 포함하는 슬라이스를 생성한다.
	var s []int = primes[1:4]
	fmt.Println(s)
}