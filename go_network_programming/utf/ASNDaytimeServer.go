package main

import (
	"encoding/asn1"
	"fmt"
	"os"
	"net"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err!= nil {
            fmt.Println("Error accepting connection:", err.Error())
            continue
        }
		daytime := time.Now()
		// Igonre return network errors
		mdata, _ := asn1.Marshal(daytime)
		conn.Write(mdata)
		conn.Close() 
	}
}

func checkError(err error) {
	if err!= nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}