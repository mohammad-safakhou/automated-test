package handlers

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
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
	PrivateKey, err = x509.ParsePKCS1PrivateKey([]byte(privateStr))
	if err != nil {
		panic(err)
	}

	psqlDb.Close()
}
