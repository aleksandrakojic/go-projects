package main

import (
    "fmt"
    "net"
    "os"
	"io"
	"bytes"
)

func main() {
	if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host port\n", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]

    conn, err := net.Dial("tcp", service)
    checkError(err)

    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))
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
        fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
        os.Exit(1)
    }
}