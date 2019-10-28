package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/duo-labs/webauthn/webauthn"

	"github.com/equinor/no-factor-auth/controllers"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	authServer string
	tenantID   string
	certPath   string
)

var (
	Version = ""
	v       = flag.Bool("version", false, "Display version")
)

func setup(e *echo.Echo) error {

	com := e.Group("/common")
	com.GET("/.well-known/openid-configuration", controllers.OidcConfig)
	com.GET("/discovery/keys", controllers.Jwks)
	com.GET("/oauth2/authorize", controllers.Authorize)
	com.GET("/oauth2/token", controllers.Token)

	web, err := webauthn.New(&webauthn.Config{
		RPDisplayName: "No Factor Auth",        // Display Name for your site
		RPID:          "localhost",             // Generally the FQDN for your site
		RPOrigin:      "http://localhost:8089", // The origin URL for WebAuthn requests
	})

	if err != nil {
		return err
	}
	wa := &controllers.WebAuthN{UserStore: nil, Authn: web}
	com.GET("/makeCredential", wa.BeginRegistration)
	com.POST("/makeCredential", wa.FinishRegistration)
	com.GET("/user", wa.BeginLogin)
	com.GET("/assertion", wa.BeginLogin)
	com.POST("/assertion", wa.FinishLogin)
	return nil
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

	authServer = os.Getenv("AUTHSERVER")
	tenantID = os.Getenv("TENANT_ID")
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	setup(e)

	e.Logger.Fatal(e.Start("0.0.0.0:8089"))
}
