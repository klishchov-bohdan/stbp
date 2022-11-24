package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}

func prefix0(otp string) string {
	if len(otp) == 6 {
		return otp
	}
	for i := (6 - len(otp)); i > 0; i-- {
		otp = "0" + otp
	}
	return otp
}

func getHOTPToken(secret string, interval int64) string {

	key, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
	check(err)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(interval))
	hash := hmac.New(sha1.New, key)
	hash.Write(bs)
	h := hash.Sum(nil)
	o := h[19] & 15
	var header uint32
	r := bytes.NewReader(h[o : o+4])
	err = binary.Read(r, binary.BigEndian, &header)
	check(err)
	h12 := (int(header) & 0x7fffffff) % 1000000
	otp := strconv.Itoa(int(h12))

	return prefix0(otp)
}

func getTOTPToken(secret string) string {

	interval := time.Now().Unix() / 30
	return getHOTPToken(secret, interval)
}

func main() {
	secret := "myawesomesecretq"
	otp := getTOTPToken(secret)

	fmt.Println("totp code:", otp)
}
