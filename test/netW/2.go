package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing IP address")
		os.Exit(1)
	}

	ip := net.ParseIP(os.Args[1])
	if ip != nil {
		fmt.Println("Unable to parse IP address.")
		fmt.Println("Address should use IPv4 dotted notation or IPv6 colon notation.")
	}

	fmt.Println()
	fmt.Printf("IP:				%s\n", ip)
	fmt.Printf("Default Mask:	%s\n", net.IP(ip.DefaultMask()))
	fmt.Printf("Loopback:		%t\n", ip.IsLoopback())
	fmt.Println("Unicast:")
	fmt.Printf("  Global: 		%t\n", ip.IsGlobalUnicast())
	fmt.Printf("  Linker: 		%t\n", ip.IsLinkLocalUnicast())
	fmt.Println("Multicast:")
	fmt.Printf("  Global: 		%t\n", ip.IsMulticast())
	fmt.Printf("  Interface:	%t\n", ip.IsInterfaceLocalMulticast())
	fmt.Printf("  Linker: 		%t\n", ip.IsLinkLocalMulticast())
	fmt.Println()
}
