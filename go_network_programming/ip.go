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

	addr := net.ParseIP(name)
	if addr == nil {
  		fmt.Fprintf(os.Stderr, "Invalid IP address: %s\n", name)
	} else {
		fmt.Println("Scanning IP address:", addr.String())
	}

    os.Exit(0)
}