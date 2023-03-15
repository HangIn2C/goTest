package main

import (
	"fmt"
	"math"
)
// Go는 클래스를 가지지 않는다.
// 하지만, 그와 같은 타입의 메소드를 정의 할 수 있다.
// 그 메소드는 특별한 receiver인자가 있는 함수이다.
// 그 receiver는 func 키워드와 메서드 이름 사이의 자체 인수 목록에 나타낸다.
// 아래의 예제에서 Abs 메소드에는 v라는 이름의 Vertex 유형의 리시버가 있다.
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // Vertex 타입의 v 리시버 // Abs라는 메소드
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}