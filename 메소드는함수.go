package main

import (
	"fmt"
	"math"
)
// 메소드는 리시버 인수가 있는 함수이다.
// 다음은 기능 변경 없이 일반 함수로 작성된 Abs 이다.
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
