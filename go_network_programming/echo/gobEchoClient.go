package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
	"io"
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
	person := Person{
		Name: Name{
            Family: "Smith",
            Personal: "John",
        },
        Email: []Email{
            Email{
                Address: "john@example.com",
                Kind:    "work",
            },
            Email{
                Address: "john.smith@example.com",
                Kind:    "home",
            },
        },
    }
	if len(os.Args)!=2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
        os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	for n := 0; n < 10; n++ {
		encoder.Encode(person)
        var newPerson Person
        decoder.Decode(&newPerson)
        fmt.Println("Received:", newPerson.String())
	}
	os.Exit(0)
}

func checkError(err error) {
	if err!= nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

    result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
        result.Write(buf[0:n])
        if err!= nil {
            if err == io.EOF {
                return result.Bytes(), nil
            }
            return nil, err
        }
		return result.Bytes(), nil
	}
}

