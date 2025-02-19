package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
	fileServer := http.FileServer(http.Dir("/var/www"))
	http.Handle("/", fileServer)

	http.HandleFunc("/cgi-bin/printenv", printEnv)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
        fmt.Println("Error starting server: ", err)
        os.Exit(1)
    }
}

func printEnv(writer http.ResponseWriter, r *http.Request) {
	env := os.Environ()
	writer.Write([]byte("<h1>Environment</h1>\n<pre>"))
	for _, v := range env {
        writer.Write([]byte(v + "\n"))
    }
	writer.Write([]byte("</pre>"))
}