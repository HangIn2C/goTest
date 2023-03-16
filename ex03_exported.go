package main

import (
	"fmt"
	"math"
)
// Export되는 이름(public : 대문자로 시작, private : 소문자로 시작)
// Go에서는 대문자로 시작하는 이름이 export 된다.
// 예를 들어 Pi가 math 패키지에서 export 되었듯이 Pizza는 export 되는 이름이다.
// pizza 와 pi 는 대문자로 시작하지 않으므로 export되지 않는다.

func main() {
	// fmt.Println(math.pi) // error
	fmt.Println(math.Pi)
}
