package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	reader := rand.Reader
	bitSize := 512
	privateKey, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	fmt.Println("private key primes", privateKey.Primes[0].String(), privateKey.Primes[1].String())
	fmt.Println("Private key exponent", privateKey.D.String())

	publicKey := privateKey.PublicKey
	fmt.Println("public key modulus", publicKey.N.String())
	fmt.Println("public key exponent", publicKey.E)

	saveGobKey("private.key", privateKey)
	saveGobKey("public.key", publicKey)

	savePEMKey("private.pem", privateKey)
}

func checkError(err error) {
	if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}

func saveGobKey(filename string, key interface{}) {
	file, err := os.Create(filename)
    checkError(err)
    defer file.Close()

    encoder := gob.NewEncoder(file)
    err = encoder.Encode(key)
    checkError(err)
}

func savePEMKey(filename string, key *rsa.PrivateKey) {
	privBytes := x509.MarshalPKCS1PrivateKey(key)
    block := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: privBytes,
    }

    file, err := os.Create(filename)
    checkError(err)
    defer file.Close()

    err = pem.Encode(file, block)
    checkError(err)
}