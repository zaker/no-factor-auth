package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/zaker/no-factor-auth/well-known"
	"net/http"
)

var routes []byte

func index(c echo.Context) error {

	return c.String(http.StatusOK, string(routes))
}

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", index)

	e.GET("/.well-known", wellknown.Index)

	var err error
	routes, err = json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Panic(err)
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
