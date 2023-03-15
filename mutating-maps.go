package main

import "fmt"

func main() {
	m := make(map[string]int)

	// m 맵에 요소를 추가하거나 업데이트 하기
	m["Answer"] = 42
	fmt.Println("The value : ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value : ", m["Answer"])

	// 요소 제거하기
	delete(m, "Answer")
	fmt.Println("The value : ", m["Answer"])

	// 2개의 값을 할당하여 키가 존재하는지 테스트 할 수 있다.
	// 만약 key가 m 안에 있다면 ok는 true, 아니라면 false
	v, ok := m["Answer"]
	fmt.Println("The value : ", v, "Present?", ok)

}
