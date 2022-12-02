package main

import (
	"io/ioutil"
	"log"
	"obfuscate/obfuscator"
	"os"
)

func main() {
	fileStr, _ := ioutil.ReadFile("./sha256/sha256.go")
	result, err := obfuscator.Obfuscate(fileStr)

	if err != nil {
		log.Fatal(err)
	}
	create, err := os.Create("./obfuscator/sha256.go")
	if err != nil {
		log.Fatal(err)
	}
	defer func(create *os.File) {
		err := create.Close()
		if err != nil {

		}
	}(create)
	_, err = create.Write([]byte(result))
	if err != nil {
		log.Fatal(err)
	}
}
