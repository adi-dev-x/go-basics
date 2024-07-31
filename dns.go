package main

import (
	"fmt"
	"net"
)

func main() {
	// Domain name to resolve
	domain := "youtube.com"

	// Resolve domain name to IP addresses
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print IP addresses
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
