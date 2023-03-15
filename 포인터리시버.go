package main

import (
	"fmt"
	"math"
)

// 포인터 리시버로 메소드를 선언 할 수 있다.
// 이는 리시버 유형이 일부 유형 T에 대한 리터럴 구문 *T를 가짐을 의미한다.
// (또한 T 자체는 *int와 같은 포인터가 될 수 없다.)
// 예를 들어, 여기서 Scale 메소드는 *Vertex에 정의되어 있다.
// 포인터 리시버가 있는 메소드는 Scale 처럼 리시버가 가리키는 값을 수정할 수 있다.
// 메소드는 종종 리시버를 수정해야하기에 포인터 리시버가 값 리시버보다 더 일반적이다.

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


// 값 리시버를 사용하면 Scale메서드가 원래 Vertex 값의 복사본에서 작동한다.
// (이것은 다른 함수 인수와 동일하다.)

// Scale 메소드에는 main 함수에 선언된 Vertex 값을 변경하기 위한 포인터 리시버가 있어야 한다.
// Scale 함수 선언에서 * 를 제거하면 값은 5, * 가 있으면 50
// 선언된 구조체의 값을 직접 바꾸기 위해서 * 를 사용한다?
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
