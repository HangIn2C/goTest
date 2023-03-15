package main

import "fmt"

type Vertex struct {
	X, Y int
}

// 구조체 리터럴은 필드 값을 나열하여 새로 할당된 구조체 값을 나타낸다.
// Name : 구문으로 필드의 하위 집합만 나열 할 수 있다.
// 특별한 접두사 &은 구조체 값으로 포인터를 반환한다.
var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y : 0 is implicit
	v3 = Vertex{}
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
