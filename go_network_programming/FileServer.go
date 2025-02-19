package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
	fileServer := http.FileServer(http.Dir("/var/www"))

	err := http.ListenAndServe(":8080", fileServer)

	if err != nil {
        fmt.Println("Error starting server: ", err)
        os.Exit(1)
    }
}