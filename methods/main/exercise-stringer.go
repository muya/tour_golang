package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr

func (ip IPAddr) String() string {
	var s string
	for i := 0; i < len(ip); i++ {
		if i == 0 {
			s = fmt.Sprintf("%d", ip[i])
		} else {
			s = fmt.Sprintf("%s.%d", s, ip[i])
		}
	}

	return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
