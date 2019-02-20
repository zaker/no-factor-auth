package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestJwks(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	if assert.NoError(t, jwksEndpoint(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		keys := jwks{}
		json.Unmarshal(rec.Body.Bytes(), &keys)
		assert.Equal(t, keys.Keys[0].Kid, "1")
		p, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(getCert("cert.pem")))
		if err != nil {
			log.Panic(err)
		}
		b64 := base64.StdEncoding.EncodeToString
		assert.Equal(t, keys.Keys[0].N, b64(p.PublicKey.N.Bytes()))
	}
}
