package main

import (
	"fmt"
	"time"
)
// goroutine은 Go 런타임에 의해 
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	go say("hello")
	say("1")
}
