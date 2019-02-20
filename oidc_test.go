package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestOidc(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	authServer = "http://some.host:6666"
	tenantID = "common"
	c := e.NewContext(req, rec)
	var oidc OpenIDConfig
	if assert.NoError(t, getOidc(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		json.Unmarshal(rec.Body.Bytes(), &oidc)
		assert.Equal(t, oidc.JwksURI, authServer+"/"+tenantID+"/discovery/keys")
	}
}
