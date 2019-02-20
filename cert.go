package main

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

func getCert(f string) string {
	s, err := ioutil.ReadFile(f)
	if err != nil {
		log.Panic(err)
		return ""
	}

	return string(s)
}

func getPubkey() *rsa.PublicKey {
	p, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(getCert("cert.pem")))
	if err != nil {
		return nil
	}

	return &p.PublicKey
}
