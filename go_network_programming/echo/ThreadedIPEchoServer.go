package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
	service := ":1200"
	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
        if err!= nil {
            fmt.Println("Error accepting connection:", err.Error())
            continue
        }

        go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

    var buf [512]byte

    for {
        n, err := conn.Read(buf[0:])
        if err!= nil {
            return
        }

        _, err2 := conn.Write(buf[0:n])
		if err2!= nil {
            fmt.Println("Error writing to connection:", err2.Error())
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
