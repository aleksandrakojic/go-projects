package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
	// Check if the required arguments are provided
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr, "Usage: %s network-type service\n", os.Args[0])
        os.Exit(1)
    }

    netwrokType := os.Args[1]
    service := os.Args[2]

	port, err := net.LookupPort(netwrokType, service)
    if err != nil {
        fmt.Println("Error: ", err.Error())
        os.Exit(2)
    }

    // Perform network scanning
    fmt.Println("Service port:", port)

    os.Exit(0)
}