package main

import (
	"aeslab/aes"
	"fmt"
	"log"
)

func main() {
	key := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
	_aes, err := aes.NewAES(key)
	if err != nil {
		log.Fatal(err)
	}
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	enc := _aes.EncryptCTR([]byte("hello"), iv)
	for _, c := range enc {
		fmt.Printf("%x ", c)
	}
	fmt.Println()
	dec := _aes.DecryptCTR(enc, iv)
	fmt.Println(string(dec))
}