package controllers

import (
	"net/http"
	"strings"

	"github.com/equinor/no-factor-auth/services"

	"github.com/labstack/echo/v4"
)

// StdOidcConfigURI is the standard andpoint for oidc config
const StdOidcConfigURI = "/.well-known/openid-configuration"

// OidcConfig returns config for host
func OidcConfig(c echo.Context) error {

	// baseUrl := c

	hostURL := "http://" + c.Request().Host + strings.TrimSuffix(c.Request().URL.String(), StdOidcConfigURI)
	oidc := services.Default()
	oidc.JwksURI = hostURL + "/discovery/keys"
	oidc.Issuer = hostURL
	oidc.AuthorizationEndpoint = hostURL + "/oauth2/authorize"
	return c.JSON(http.StatusOK, oidc)
}
