package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
        fmt.Println("Usage: ", os.Args[0], "https://host:port/page https://host:port/page")
        os.Exit(1)
    }

    url, err := url.Parse(os.Args[1])
    checkError(err)

	if url.Scheme != "https" {
		fmt.Println("Invalid URL scheme. Only https is supported.")
		os.Exit(1)
	}

	transport := &http.Transport{}
    transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    client := &http.Client{Transport: transport}

    request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)

	response, err := client.Do(request)
	checkError(err)

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
        os.Exit(2)
	}
	fmt.Println("Response ok")

	chSet := getCharset(response)
	if chSet != "utf-8" {
        fmt.Println("Cannot handle", chSet)
		os.Exit(4)
    }

	var buf [512]byte
	reader := response.Body
	fmt.Println("The response body is")

	for {
		n, err := reader.Read(buf[0:])
        if err != nil {
            os.Exit(0)
        }
        fmt.Print(string(buf[0:n]))
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
    if contentType == "" {
        return "UTF-8"
    }
	idx := strings.Index(contentType, "charset:")
	if idx == -1 {
        return "UTF-8"
    }
    return strings.Trim(contentType[idx:], " ")
}