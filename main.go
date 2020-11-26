package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cloudflare/cfssl/revoke"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("path to certificate must be passed as an argument")
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pemBlock, _ := pem.Decode(bytes)
	if pemBlock == nil {
		fmt.Println("no PEM block found")
		os.Exit(1)
	}

	cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	revoked, ok, err := revoke.VerifyCertificateError(cert)
	fmt.Printf("revoked: %v, ok: %v, err: %v\n", revoked, ok, err)
}
