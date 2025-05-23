package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type Person struct {
    Name Name
    Email []Email
}

type Name struct {
    Family string
    Personal  string
}

type Email struct {
    Address string
    Kind    string
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
    for _, v := range p.Email {
        s +=  "\n" + v.Kind + ": " + v.Address
    }
    return s
}

func main() {
	service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
        if err!= nil {
            continue
        }

		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		for n := 0; n < 10; n++ {
			var person Person
			decoder.Decode(&person)
			fmt.Println("Received:", person.String())
			encoder.Encode(person)
		}
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}