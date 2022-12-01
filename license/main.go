package main

import (
	"encoding/json"
	"fmt"
	"lic/keygen"
	"log"
	"time"
)

type MyLicense struct {
	Email string    `json:"email"`
	End   time.Time `json:"end"`
}

func main() {
	privateKey, err := keygen.NewPrivateKey()
	if err != nil {
		log.Fatal(err)

	}

	doc := MyLicense{
		"limbo@test.com",
		time.Now().Add(time.Hour * 24 * 365), // 1 year
	}

	docBytes, err := json.Marshal(doc)
	if err != nil {
		log.Fatal(err)

	}

	license, err := keygen.NewLicense(privateKey, docBytes)
	if err != nil {
		log.Fatal(err)

	}

	str64, err := license.ToB64String()
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("license key:", str64)

	// validate
	publicKey := privateKey.GetPublicKey()

	if ok, err := license.Verify(publicKey); err != nil {
		log.Fatal(err)
	} else if !ok {
		log.Fatal("Invalid license signature")
	}

	res := MyLicense{}
	if err := json.Unmarshal(license.Data, &res); err != nil {
		log.Fatal(err)
	} else if res.End.Before(time.Now()) {
		log.Fatalf("License expired on: %s", res.End.String())
	} else {
		fmt.Printf("Licensed to %s until %s \n", res.Email, res.End.Format("2006-01-02"))
	}
}
