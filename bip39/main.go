package main

import (
	"bip39/bip39"
	"fmt"
	"github.com/tyler-smith/go-bip32"
)

func main() {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")
	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)

	departmentKeys := map[string]*bip32.Key{}
	departmentKeys["Sales"], _ = masterKey.NewChildKey(0)
	departmentKeys["Marketing"], _ = masterKey.NewChildKey(1)
	departmentKeys["Engineering"], _ = masterKey.NewChildKey(2)
	departmentKeys["Customer Support"], _ = masterKey.NewChildKey(3)

	departmentAuditKeys := map[string]*bip32.Key{}
	departmentAuditKeys["Sales"] = departmentKeys["Sales"].PublicKey()
	departmentAuditKeys["Marketing"] = departmentKeys["Marketing"].PublicKey()
	departmentAuditKeys["Engineering"] = departmentKeys["Engineering"].PublicKey()
	departmentAuditKeys["Customer Support"] = departmentKeys["Customer Support"].PublicKey()

	for department, pubKey := range departmentAuditKeys {
		fmt.Println(department, pubKey)
	}
}
