package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/equinor/no-factor-auth/services"

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
	var oidc services.OpenIDConfig
	if assert.NoError(t, OidcConfig(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		json.Unmarshal(rec.Body.Bytes(), &oidc)
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
	var oidc services.OpenIDConfig
	if assert.NoError(t, OidcConfig(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		json.Unmarshal(rec.Body.Bytes(), &oidc)
		assert.Equal(t, authServer+"/discovery/keys", oidc.JwksURI)
	}
}

// func TestOidcConfig(t *testing.T) {
// 	type args struct {
// 		c echo.Context
// 	}
// 	args.c = e.NewContext(req, rec)
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"Returns oidc-configuration",}
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := OidcConfig(tt.args.c); (err != nil) != tt.wantErr {
// 				t.Errorf("OidcConfig() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
