package main

import "fmt"

// channel은 buffered 될 수 있다.
// (buffered channel이란 channel이 버퍼를 가질 수 있다는 의미)
// buffered channel을 초기화하기 위해 make에 두 번째 인자로 buffer를 제공
// ch := make(chan int, 100)
// buffered channel로의 전송은 그 buffer의 사이즈가 꽉 찼을 때에만 block 된다.
// buffer로 부터의 수신은 그 buffer가 비어있을 때 block 된다.
func main() {
	ch := make(chan int, 2) // 버퍼가 2개 3개 째에서 에러가 발생
	ch <- 1
	ch <- 2
	// ch <- 3
	fmt.Println(<-ch) // 버퍼 1
	fmt.Println(<-ch) // 버퍼 2
	// fmt.Println(<-ch)
}
