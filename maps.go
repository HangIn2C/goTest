package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 맵은 키를 값에 매핑한다.
// 맵의 초기값은 nil이며 nil 맵은 키도 없고, 키를 추가할 수도 없다.
// make 함수는 주어진 타입의 초기화되고 사용 준비가 된 맵을 반환한다.
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
