package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Go에서 switch-case는 break를 생략할 수 있다.
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os { // runtime.GOOS : os 실행 환경 = onwindows
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
