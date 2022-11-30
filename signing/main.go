package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"signing/rsa"
)

func signFile(priKey string, filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	hash := sha256.Sum256(file)
	signature, err := rsa.PriKeyEncrypt(string(hash[:]), priKey)
	if err != nil {
		return err
	}
	created, err := os.Create("./sha.signature")
	if err != nil {
		return err
	}
	_, err = created.Write([]byte(signature))
	if err != nil {
		return err
	}
	return nil
}

func verifySignature(pubKey string, signature string, filePath string) (bool, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, err
	}
	hash := sha256.Sum256(file)
	decrypt, err := rsa.PublicDecrypt(signature, pubKey)
	if err != nil {
		return false, err
	}
	if decrypt != string(hash[:]) {
		return false, nil
	}
	return true, nil
}

func main() {
	priKey, err := ioutil.ReadFile("./keys/key.pem")
	if err != nil {
		log.Fatal(err)
	}
	pubKey, err := ioutil.ReadFile("./keys/key.pub")
	if err != nil {
		log.Fatal(err)
	}
	err = signFile(string(priKey), "./sha")
	if err != nil {
		log.Fatal(err)
	}
	signature, err := ioutil.ReadFile("./sha.signature")
	if err != nil {
		log.Fatal(err)
	}
	isValid, err := verifySignature(string(pubKey), string(signature), "./sha")
	if err != nil {
		return
	}
	if isValid {
		fmt.Println("the program is valid")
	} else {
		fmt.Println("the program invalid")
	}
}
