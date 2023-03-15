package main

import (
	"fmt"
	"math"
)

// 이와 동등한 일은 역방향에서 일어날 수 있다.
// 값 인수를 사용하는 함수는 다음과 같은 특정 유형의 값을 사용해야 한다.

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// value receiver가 있는 메소드는 다음과 같이 호출될 때, 값이나 포인터를 리시버로 사용한다.
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))
	// fmt.Println(AbsFunc(&v)) // compile error

	// var v Vertex
	// p := &v
	// fmt.Println(p.Abs())
	// p.Abs() 라는 메서드는 (*p).Abs()로 해석된다.
	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
