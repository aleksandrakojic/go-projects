package main

import (
    "fmt"
    "net"
    "os"
	"encoding/json"
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
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
        os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	for n := 0; n < 10; n++ {
		encoder.Encode(person)
        var newPerson Person
        err = decoder.Decode(&newPerson)
        checkError(err)
        fmt.Println("Received person:", newPerson.String())
	}
	os.Exit(0)
}

func checkError(err error) {
	if err!= nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
