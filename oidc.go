package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func defaultOidc() *OpenIDConfig {
	oidc := &OpenIDConfig{
		// JwksURI:               authServer + "/" + tenantID + "/discovery/keys",
		SubjectTypesSupported:               []string{"pairwise"},
		IDTokenEncryptionAlgValuesSupported: []string{"RS256"},
		IDTokenSigningAlgValuesSupported:    []string{"RS256"},
	}

	return oidc
}

func getOidc(c echo.Context) error {

	oidc := defaultOidc()
	oidc.JwksURI = authServer + "/" + tenantID + "/discovery/keys"
	oidc.Issuer = authServer
	oidc.AuthorizationEndpoint = authServer + "/" + tenantID + "/oauth2/authorize"
	return c.JSON(http.StatusOK, oidc)
}
