package main

import (
	"encoding/base64"
	"math/big"
	"net/http"

	"github.com/labstack/echo"
)

type jwk struct {
	Kty string   `json:"kty"`
	Use string   `json:"use"`
	Kid string   `json:"kid"`
	X5T string   `json:"x5t"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5C []string `json:"x5c"`
}

type jwks struct {
	Keys []jwk `json:"keys"`
}

func getKeyset() (*jwks, error) {

	pub := getPubkey()
	b64 := base64.StdEncoding.EncodeToString

	e := big.Int{}
	e.SetUint64(uint64(pub.E))

	keys := jwks{
		Keys: []jwk{
			{
				Kty: "RSA",
				N:   b64(pub.N.Bytes()),
				E:   b64(e.Bytes()),
				Kid: "1",
				X5T: "1",
			}}}
	return &keys, nil
}

func jwksEndpoint(c echo.Context) error {

	jwks, err := getKeyset()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, jwks)
}
