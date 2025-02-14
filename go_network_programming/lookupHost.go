package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Check if the required arguments are provided
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
        os.Exit(1)
    }

    // Get the target host from the command-line argument
    name := os.Args[1]

    // Resolve the hostname to an IP address
    addrs, err := net.LookupIP(name)
    if err != nil {
        fmt.Printf("Error resolving hostname: %v\n", err)
        os.Exit(2)
    }

    // Print the IP addresses
    fmt.Printf("IP addresses for %s:\n", name)
    for _, addr := range addrs {
        fmt.Println(addr)
    }
	os.Exit(0)
}