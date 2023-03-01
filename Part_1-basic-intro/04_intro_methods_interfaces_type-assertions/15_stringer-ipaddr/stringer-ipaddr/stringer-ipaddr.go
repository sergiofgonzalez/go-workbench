package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

// func (ipAddr IPAddr) String() string {
// 	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
// }

func (ipAddr IPAddr) String() string {

	var ipAddrAsStr [4]string
	for i, ipComp := range ipAddr {
		ipAddrAsStr[i] = string(fmt.Sprintf("%v", ipComp))
	}
	return strings.Join(ipAddrAsStr[:], ".")
}

func main() {
	myIpAddress := IPAddr{34, 103, 129, 15}
	fmt.Println(myIpAddress)

	hosts := map[string]IPAddr{
		"mine": {34, 103, 129, 15},
		"localhost": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}