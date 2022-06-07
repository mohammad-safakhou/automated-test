package main

//go:generate sqlboiler --wipe --no-tests psql -o usecase_models/boiler

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"test-manager/cmd"
	"test-manager/repos"
	"test-manager/utils"
)

func main() {
	//privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	//if err != nil {
	//	return
	//}
	//
	//privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	//privkey_pem := pem.EncodeToMemory(
	//	&pem.Block{
	//		Type:  "RSA PRIVATE KEY",
	//		Bytes: privkey_bytes,
	//	},
	//)
	//pubk := pem.EncodeToMemory(
	//	&pem.Block{
	//		Type:  "RSA PUBLIC KEY",
	//		Bytes: x509.MarshalPKCS1PublicKey(&privkey.PublicKey),
	//	},
	//)
	//fmt.Println(string(privkey_pem))
	//fmt.Println(string(pubk))
	//fmt.Printf("%q\n", string(x509.MarshalPKCS1PrivateKey(privateKey)))
	//x509.ParsePKCS1PrivateKey()
	//utils.RsaOaepEncrypt(privateKey.)
	psqlDb, err := utils.PostgresConnection("localhost", "5432", "root", "root", "tester", "disable")
	if err != nil {
		panic(err)
	}

	authInfoRepo := repos.NewAuthInfoRepositoryRepository(psqlDb)
	privateStr, err := authInfoRepo.GetAuthKeys(context.TODO())
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode([]byte(privateStr))
	if block == nil {
		panic("failed to parse PEM block containing the key")
	}
	PrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	s, err := utils.RsaOaepEncrypt("mypassword2", PrivateKey.PublicKey)
	fmt.Println(s)
	cmd.Execute()
}
