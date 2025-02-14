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
	saveGob("person.gob", person)
}

func saveGob(filename string, data interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)

	encoder := gob.NewEncoder(outFile)

	err = encoder.Encode(data)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err!= nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

