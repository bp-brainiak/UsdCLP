package main

import (
	"backend/config"
	"backend/yahooQuery"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// Routes
	e.GET("/usdSpot", GetDolarSpot)

	// Start server
	e.Logger.Fatal(e.Start(config.GetValue("backport")))
}

func GetDolarSpot(c echo.Context) error {
	res := yahooQuery.RetornaSpot()
	log.Println(res)
	return c.JSON(http.StatusOK, res)
}
