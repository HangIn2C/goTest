package main

import (
	"fmt"
	"math"
	"math/rand"
)
// Package
// 모든 Go 프로그램은 패키지로 구성된다.
// 프로그램은 main 패키지에서 실행을 시작한다.
// 관습적으로, 패키지의 이름은 import 경로의 마지막 요소와 같다.
// 예를 들어 "math/rand" 패키지는 package rand문으로 시작하는 파일들로 구성된다.

// import
// 괄호로 import를 그룹 지어진것을 "factored" import문이라 한다.
// import 문은 아래와 같이 나누어 적을수 있다. 
// import "fmt"
// import "math"
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}
