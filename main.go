package main

import (
	"log"
	"os"

	"github.com/equinor/no-factor-auth/controllers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	authServer string
	tenantID   string
	certPath   string
)

func setup(e *echo.Echo) {

	com := e.Group("/common")
	com.GET("/.well-known/openid-configuration", controllers.OidcConfig)
	com.GET("/discovery/keys", controllers.Jwks)
	com.GET("/oauth2/authorize", controllers.Authorize)
}

func main() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file", err)
	}

	authServer = os.Getenv("AUTHSERVER")
	tenantID = os.Getenv("TENANT_ID")
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	setup(e)

	e.Logger.Fatal(e.Start("0.0.0.0:8089"))
}
