package main

import (
    "fmt"
    "net"
    "os"
	"time"
	"encoding/asn1"
	"bytes"
	"io"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
        os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	var newtime time.Time
	_, err = asn1.Unmarshal(result, &newtime)
	checkError(err)

	fmt.Println("After marshal/unmarshal", newtime.String())

	os.Exit(0)
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

    result := bytes.NewBuffer(nil)
    var buf [512]byte

    for {
        n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
        if err != nil {
			if err == io.EOF {
                break
            }
			return nil, err
        }
    }
    return result.Bytes(), nil
}

func checkError(err error) {
	if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
