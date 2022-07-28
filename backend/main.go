package main

import (
	"backend/config"
	"backend/yahooQuery"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.GET("/chart", GetChartData)
	// Start server
	e.Logger.Fatal(e.Start(config.GetValue("backport")))
}

func GetDolarSpot(c echo.Context) error {
	res := yahooQuery.RetornaSpot()

	return c.JSON(http.StatusOK, res)
}

func GetChartData(c echo.Context) error {
	fmt.Println("-------------------------------------------------------")

	res := yahooQuery.GetChart(
		c.QueryParam("range"),
		c.QueryParam("region"),
		c.QueryParam("interval"),
		c.QueryParam("lang"),
	)

	return c.JSON(http.StatusOK, res)
}
