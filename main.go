package main

//go:generate sqlboiler --wipe --no-tests psql -o usecase_models/boiler

import "test-manager/cmd"

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
	cmd.Execute()
}
