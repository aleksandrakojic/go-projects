package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.AcceptTCP()
        if err!= nil {
            fmt.Println("Error accepting connection:", err.Error())
            continue
        }

        go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
        if err!= nil {
            return
        }

        fmt.Println("Received:", string(buf[0:]))

        _, err = conn.Write(buf[0:n])        
        if err!= nil {
            return
        }
	}
}

func checkError(err error) {
	if err!= nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}