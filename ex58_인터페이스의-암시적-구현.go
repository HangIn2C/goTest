package main

import (
	"fmt"
)
// type implements는 메소드를 실행함으로써 인터페이스를 구현한다.
// 명시적 intent의 선언도, "implemnetation"의 키워드도 없다.

// 암시적 인터페이스는 인터페이스의 정의를 구현으로부터 분리하며, 이는 사전 정렬 없이
// 어떠한 패키지에 등장할 수 있다.
type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
