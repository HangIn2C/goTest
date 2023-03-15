package main

import "fmt"

// 반환형이 func() int인 함수형을 선언
func fibonacci() func() int {
	fn := 0
	fn1 := 0
	return func() int {
		fn2 := fn1 + fn
		fn = fn1
		fn1 = fn2
		if fn1 == 0 {
			fn++
		}
		return fn2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
