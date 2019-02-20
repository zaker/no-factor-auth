package main

import (
	"github.com/labstack/echo/v4"
)

func setup(e *echo.Echo) {

	com := e.Group("/common")
	com.GET("/.well-known/openid-configuration", getOidc)
	com.GET("/discovery/keys", jwksEndpoint)
	com.GET("/oauth2/authorize", authorize)
}
