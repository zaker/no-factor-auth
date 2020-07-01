package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/equinor/no-factor-auth/controllers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	Version = ""
	v       = flag.Bool("version", false, "Display version")
)

func setup(e *echo.Echo) {

	com := e.Group("/common")
	com.GET("/.well-known/openid-configuration", controllers.OidcConfig)
	com.GET("/discovery/keys", controllers.Jwks)
	com.GET("/oauth2/authorize", controllers.Authorize)
	com.GET("/oauth2/token", controllers.Token)
}

func version() {
	fmt.Println("Version:", Version)
	fmt.Println("Go Version:", runtime.Version())
	os.Exit(0)
}

func main() {

	flag.Parse()

	if *v {
		version()
	}

	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	setup(e)

	e.Logger.Fatal(e.Start("0.0.0.0:8089"))
}
