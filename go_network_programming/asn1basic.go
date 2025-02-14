package main

import (
	"encoding/asn1"
    "fmt"
    "time"
    "os"
)

func main() {
	currentTime := time.Now()

    asn1Time, err := asn1.Marshal(currentTime)
    checkError(err)
	fmt.Printf("ASN.1 encoded time: %s\n", string(asn1Time))

    var newtime = new(time.Time)
    _, err = asn1.Unmarshal(asn1Time, newtime)
    checkError(err)

	fmt.Println("After marshal/unmarshal: ", newtime.String())
	s := "hello \u00bc"
	fmt.Println("Before marshalling: ", s)

	mdata2, err := asn1.Marshal(s)
	checkError(err)
	fmt.Println("Marshalled ok")

	var newstr string
	_, err2 := asn1.Unmarshal(mdata2, &newstr)
	checkError(err2)

	fmt.Println("After marshal/unmarshal: ", newstr)
}

func checkError(err error) {
	if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}