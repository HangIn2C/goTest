package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// for에서 range는 슬라이스 또는 맵의 요소들을 순회한다.
	// 슬라이스에서 range를 사용하면, 각 순회마다 두 개의 값이 반환된다.
	// 첫 번째는 인덱스이고, 두 번째는 해당 인덱스 값의 복사본입니다.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
