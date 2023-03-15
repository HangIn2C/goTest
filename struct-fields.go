package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	// 구조체의 필드는 .(dot)으로 접근할수 있다.
	fmt.Println(v.X)
}
