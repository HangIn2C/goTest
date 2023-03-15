package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 최상위 타입이 타입 이름일 겨우, 리터럴의 요소에서 생략 가능
var m = map[string]Vertex{
	"Bell Labs": {40.68, -74.399},
	"Google":    {37.422, -122.084},
}

func main() {
	fmt.Println(m)
}
