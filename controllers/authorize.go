package controllers

import (
	"net/http"
	"net/url"
	"time"

	"github.com/equinor/no-factor-auth/config"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo/v4"
)

func newToken(sub, iss, aud, nonce, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":       sub,
		"nbf":       time.Now().Unix(),
		"iss":       iss,
		"aud":       aud,
		"nonce":     nonce,
		"auth_time": time.Now().Unix(),
		"acr":       "no-factor",
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(1 * time.Hour).Unix(),
		"name":      name,
	})

	token.Header = map[string]interface{}{
		"typ": "JWT",
		"alg": jwt.SigningMethodRS256.Name,
		"kid": "1",
	}

	// Sign and get the complete encoded token as a string using the secret

	tokenString, err := token.SignedString(config.PrivateKey())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Authorize provides id_token and access_token to anyone who asks
func Authorize(c echo.Context) error {
	redirectURI := c.QueryParam("redirect_uri")
	if redirectURI == "" {
		redirectURI = "/"
	}
	clientID := c.QueryParam("client_id")
	state := c.QueryParam("state")

	sub := c.QueryParam("sub")
	if len(sub) == 0 {
		sub = "anon1"
	}

	user := c.QueryParam("user")
	if len(user) == 0 {
		user = "Jane Doe"
	}

	// Sign and get the complete encoded token as a string using the secret

	tokenString, err := newToken(sub, c.Request().Host, clientID, c.QueryParam("nonce"), user)
	if err != nil {
		return err
	}
	params := url.Values{}
	params.Set("id_token", tokenString)
	params.Set("access_token", tokenString)
	params.Set("state", state)
	p := params.Encode()
	return c.Redirect(http.StatusFound, redirectURI+"?"+p+"#"+p)
}
