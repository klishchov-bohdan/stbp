package main

import (
	sha256pkg "crypto/sha256"
	"fmt"
	"sha/sha256"
)

func main() {
	msg := []byte("hello world")
	hashed1 := sha256.Hash(msg)
	hashed2 := sha256pkg.Sum256(msg)
	if hashed2 != hashed1 {
		fmt.Println("Hash codes is not matched")
	} else {
		fmt.Println("Hash codes is matched")
	}
	fmt.Println(hashed1)
	fmt.Println(hashed2)
}
