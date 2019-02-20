package controllers
import (
	"encoding/base64"
	"math/big"
	"net/http"

	"github.com/equinor/no-factor-auth/config"
	"github.com/equinor/no-factor-auth/oidc"
	"github.com/labstack/echo/v4"
)

func rsaKeyset() (*oidc.JWKS, error) {

	pub := config.PublicKey()
	b64 := base64.StdEncoding.EncodeToString

	e := big.Int{}
	e.SetUint64(uint64(pub.E))

	keys := oidc.JWKS{
		Keys: []oidc.JWK{
			{
				Kty: "RSA",
				N:   b64(pub.N.Bytes()),
				E:   b64(e.Bytes()),
				Kid: "1",
				X5T: "1",
			}}}
	return &keys, nil
}

func Jwks(c echo.Context) error {

	jwks, err := rsaKeyset()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, jwks)
}
