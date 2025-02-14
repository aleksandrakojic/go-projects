package main

import (
    "encoding/json"
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
	saveJSON("peson.json", person)
}

func saveJSON(filename string, data interface{}) error {
	outFile, err := os.Create(filename)
	checkError(err)

	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(data)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
