package main

import (
    "encoding/gob"
    "fmt"
    "os"
)

type Person struct {
    Name Name
    Email []Email
}

type Name struct {
    Family string
    Personal string
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
	var person Person
	loadGob("person.gob", &person)

	fmt.Println("Person:", person.String())
}

func loadGob(filename string, data interface{}) {
	inFile, err := os.Open(filename)
    checkError(err)

    decoder := gob.NewDecoder(inFile)

    err = decoder.Decode(data)
    checkError(err)
    inFile.Close()
}

func checkError(err error) {
	if err!= nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}