package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:post")
		os.Exit(1)
	}
	url := os.Args[1]
	response, err := http.Head(url)
	if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(2)
    }
	fmt.Println("HTTP Status:", response.Status)
	for k, v := range response.Header {
		fmt.Printf("%s: %s\n", k, v)
	}

	os.Exit(0)
}
