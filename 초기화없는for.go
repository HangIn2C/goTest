package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	// for {
	//	 무한 루프
	// }

	fmt.Println(sum)
}
