package main

import "fmt"

import "net"

func Increment(ip net.IP) {
		//fmt.Println(ip.String())
	
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		//fmt.Println(ip.String())
		//only add to the next byte if we overflowed
		if ip[i] != 0 {
			break
		}
	}
}

func main() {
	ipaddr := net.ParseIP("192.168.1.2")
  for i :=0; i <300; i++{
	  Increment(ipaddr)
	  fmt.Println("ip addr ", ipaddr)	
	}
}
