package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 맵 리터럴은 구조체 리터럴과 같지만, 키가 필요하다.
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68, -74.399,
	},
	"Google": Vertex{
		37.422, -122.084,
	},
}

func main() {
	fmt.Println(m)
}
