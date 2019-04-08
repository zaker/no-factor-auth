package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TokenOKResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	ExpiresOn    string `json:"expires_on"`
	Resource     string `json:"resource"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
}

type TokenErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	Timestamp        string `json:"timestamp"`
	TraceID          string `json:"trace_id"`
	CorrelationID    string `json:"correlation_id"`
}

// Token provides id_token and access_token to anyone who asks
func Token(c echo.Context) error {
	redirectURI := c.QueryParam("redirect_uri")
	if redirectURI == "" {
		return c.JSON(http.StatusBadRequest, TokenErrorResponse{Error: "No redirect_uri"})

	}
	clientID := c.QueryParam("client_id")
	if clientID == "" {
		return c.JSON(http.StatusBadRequest, TokenErrorResponse{Error: "No client_id"})
	}
	grantType := c.QueryParam("grant_type")
	if grantType == "" {
		return c.JSON(http.StatusBadRequest, TokenErrorResponse{Error: "No grant_type"})
	}
	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, TokenErrorResponse{Error: "No code"})
	}
	clientSecret := c.QueryParam("client_secret")
	if clientSecret == "" {
		return c.JSON(http.StatusBadRequest, TokenErrorResponse{Error: "No client_secret"})
	}
	a, err := newToken("anon1", c.Request().Host, clientID, "Foo", "Jane Doe")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, TokenOKResponse{AccessToken: a, IDToken: a, TokenType: "Bearer"})
}
