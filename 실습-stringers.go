package main

import (
	"fmt"
)

type IPAddr [4]byte

// func (ip IPAddr) split() string {
// 	return fmt.Sprintf("%v.%v.%v.%v", ip.)
// }

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v : %v\n", name, ip)
	}
}
