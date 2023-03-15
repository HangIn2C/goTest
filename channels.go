package main

import "fmt"

// channel은 당신이 채널 연산자인 <-을 통해 값을 주고 받을 수 있는 하나의 분리된 통로이다.
// ch <- v // 채널 ch에 v를 전송한다.
// v := <-ch // ch로 부터 값을 받고, 값을 v에 대입한다.
// (데이터는 화살표의 방향대로 흐른다.)
// channel은 map과 slice처럼 사용하기 전에 생성되어야만 한다.
// ch := make(chan int)
// 기본적으로 전송과 수신은 다른 한 쪽이 준비될 때까지 block 상태이다.
// 이는 명시적으로 lock이나 조건 변수 없이 goroutine이 sychronous하게 작업될 수 있도록 한다.
// 예제 코드는 두 개의 goroutine에 작업을 분산시키면서 slice 있는 숫자들을 더한다.
// 두 goroutine이 그들의 연산을 완료하면, 최종 결과를 계산한다.
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 9, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
