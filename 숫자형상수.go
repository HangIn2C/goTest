package main

import "fmt"

const (
	// 1 비트를 100번 왼쪽으로 쉬프트
	Big = 1 << 100
	// Big에 있는 값을 99번 오른쪽으로 쉬프트 -> Small = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1 // 2 * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1 // 2 * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
