package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	random := rand.Reader

	var key rsa.PrivateKey
	loadKey("private.key", key)

	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000)
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
        Subject: pkix.Name{
            Organization:  []string{"Acme Co"},
            Country:       []string{"US"},
            Province:      []string{"California"},
            Locality:      []string{"San Francisco"},
            StreetAddress: []string{"1355 Market St"},
            PostalCode:    []string{"94103"},
        },
        NotBefore:             now,
        NotAfter:              then,
        KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
        ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
        BasicConstraintsValid: true,
		IsCA: true,
		DNSNames: []string{"jan.newmarch.name", "localhost"},
    }
	derBytes, err := x509.CreateCertificate(random, &template, &template, &key.PublicKey, &key)
	checkError(err)

	certCerFile, err := os.Create("jan.newmarch.name.cer")
	checkError(err)

	certCerFile.Write(derBytes)
	certCerFile.Close()

	certPEMFile, err := os.Create("jan.newmarch.name.pem")
	checkError(err)
	pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certPEMFile.Close()

	keyPEMFile, err := os.Create("private.pem")
	checkError(err)
	pem.Encode(keyPEMFile, &pem.Block{Type: "RSA PRIVATE KEY",
	Bytes: x509.MarshalPKCS1PrivateKey(&key)})
	keyPEMFile.Close()
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
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}