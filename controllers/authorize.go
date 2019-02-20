package controllers

import (
	"net/http"
	"net/url"
	"time"

	"github.com/equinor/no-factor-auth/config"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo/v4"
)

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
		user = "Bodil Rotevatn"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":       sub,
		"nbf":       time.Now().Unix(),
		"iss":       c.Request().Host,
		"aud":       clientID,
		"nonce":     c.QueryParam("nonce"),
		"auth_time": time.Now().Unix(),
		"acr":       "no-factor",
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(1 * time.Hour).Unix(),
		"name":      user,
	})

	token.Header = map[string]interface{}{
		"typ": "JWT",
		"alg": jwt.SigningMethodRS256.Name,
		"kid": "1",
	}

	// Sign and get the complete encoded token as a string using the secret

	tokenString, err := token.SignedString(config.PrivateKey())
	if err != nil {
		return err
	}
	params := url.Values{}
	params.Set("id_token", tokenString)
	params.Set("state", state)

	return c.Redirect(http.StatusFound, redirectURI+"#"+params.Encode())
}
