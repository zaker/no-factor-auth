package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/equinor/no-factor-auth/oidc"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestOidc(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, StdOidcConfigURI, nil)
	rec := httptest.NewRecorder()
	authServer := "http://example.com"
	// tenantID := "common"
	c := e.NewContext(req, rec)
	var oidc oidc.OpenIDConfig
	if assert.NoError(t, OidcConfig(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		err := json.Unmarshal(rec.Body.Bytes(), &oidc)
		if err != nil {
			assert.NoError(t, err)
		}
		assert.Equal(t, authServer+"/discovery/keys", oidc.JwksURI)
	}
}
func TestOidcTenant(t *testing.T) {
	e := echo.New()
	tenant := "/tenant"
	req := httptest.NewRequest(http.MethodGet, tenant+StdOidcConfigURI, nil)
	rec := httptest.NewRecorder()
	authServer := "http://example.com" + tenant
	// tenantID := "common"
	c := e.NewContext(req, rec)
	var oidc oidc.OpenIDConfig
	if assert.NoError(t, OidcConfig(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		err := json.Unmarshal(rec.Body.Bytes(), &oidc)
		if err != nil {
			assert.NoError(t, err)
		}
		assert.Equal(t, authServer+"/discovery/keys", oidc.JwksURI)
	}
}
