package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"os"
)

var (
	cidr string
)

func init() {
	flag.StringVar(&cidr, "c", "", "the CIDR address")
}

func main() {
	flag.Parse()
	if cidr == "" {
		fmt.Printf("CIDR address missing")
		os.Exit(1)
	}

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Printf("failed parsing CIDR address: %v", err)
		os.Exit(1)
	}

	ones, totalBits := ipnet.Mask.Size()
	size := totalBits - ones                 // usable bits
	totalHosts := math.Pow(2, float64(size)) //2^size
	wildcardsIP := wildcard(net.IP(ipnet.Mask))
	last := lastIP(ip, net.IPMask(wildcardsIP))

	fmt.Println()
	fmt.Printf("CIDR: %s\n", cidr)
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("CIDR Block:					%s\n", cidr)
	fmt.Printf("Network:		 			%s\n", ipnet.IP)
	fmt.Printf("IP Range:					%s - %s\n", ip, last)
	fmt.Printf("Total Hosts:				%s\n", totalHosts)
	fmt.Printf("Wildcard Mask:				%s\n", wildcardsIP)
	fmt.Println()
}

func wildcard(mask net.IP) net.IP {
	var ipValue net.IP
	for _, octet := range mask {
		ipValue = append(ipValue, ^octet)
	}
	return ipValue
}

func lastIP(ip net.IP, mask net.IPMask) net.IP {
	ipIn := ip.To4() // is it an IPv4
	if ipIn == nil {
		ipIn = ip.To16() //is it an IPv6
		if ipIn == nil {
			return nil
		}
	}

	var ipVal net.IP
	for i, octet := range ipIn {
		ipVal = append(ipVal, octet|mask[i])
	}
	return ipVal
}
