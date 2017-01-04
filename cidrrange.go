package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ip, ipnet, err := net.ParseCIDR("62.76.47.0/24")
	if err != nil {
		log.Fatal(err)
	}
	endip := ip
	fmt.Println("start ip ", ip)

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		if ipnet.Contains(ip) {
			endip = ip
		}
	}
	fmt.Println("endip ", endip)
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
