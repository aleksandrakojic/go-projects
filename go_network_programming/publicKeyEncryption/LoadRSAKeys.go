package main

import (
    "crypto/rsa"
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	var key rsa.PrivateKey
	loadKey("private.key", &key)

	fmt.Println("private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("private key exponent", key.D.String())

	publicKey := key.PublicKey
	loadKey("public.key", &publicKey)

	fmt.Println("public key modulus", publicKey.N.String())
	fmt.Println("public key exponent", publicKey.E)
}

func loadKey(filename string, key interface{}) {
	file, err := os.Open(filename)
    checkError(err)
    defer file.Close()

    decoder := gob.NewDecoder(file)
    err = decoder.Decode(key)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}