package main

import (
    "fmt"
	"io"
    "net"
    "os"
)

func main() {
	if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
        os.Exit(1)
    }

    // Get the target host from the command-line argument
    service := os.Args[1]

    // Resolve the IP address to a hostname
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := io.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
        os.Exit(1)
    }
}