package main

//go:generate sqlboiler --wipe --no-tests psql -o usecase_models/boiler

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}

	publicKey := privateKey.PublicKey

	x509.MarshalPKCS1PrivateKey(privateKey)
	x509.ParsePKCS1PrivateKey()
	fmt.Println(publicKey)
	//utils.RsaOaepEncrypt(privateKey.)
	//cmd.Execute()
}
