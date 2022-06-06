package handlers

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"test-manager/repos"
	"test-manager/utils"
)

var PrivateKey *rsa.PrivateKey

func init() {
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
	PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	psqlDb.Close()
}
