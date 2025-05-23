package main

import (
	"fmt"
    "net"
    "os"
)

const (
	DIR = "DIR"
	CD = "CD"
	PWD = "PWD"
)

func main() {
	service := "0.0.0.0:1202"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
        if err!= nil {
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
			conn.Close()
            return
        }

        s := string(buf[0:n])

		if s[0:2] == CD {
			chdir(conn, s[3:])
		} else if s[0:3] == DIR {
            dirList(conn)
        } else if s[0:3] == PWD {
            pwd(conn)
        }
    }
}

func chdir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
        conn.Write([]byte("ERROR"))
    }
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(s))
		return
	}
}

func dirList(conn net.Conn) {
	// send a blank line on termination
	defer conn.Write([]byte("\n"))

    dir, err := os.Open(".")
    if err!= nil {
        return
    }

	names, err := dir.Readdirnames(-1)
	if err != nil {
        return
    }

    for _, f := range names {
        conn.Write([]byte(f + "\r\n"))
    }
}

func checkError(err error) {
	if err!= nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}