package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Check if the required arguments are provided
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
    }

    // Get the target host from the command-line argument
    name := os.Args[1]

    addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}

    fmt.Println("Resolved address is ", addr.String())
 	os.Exit(0)
}