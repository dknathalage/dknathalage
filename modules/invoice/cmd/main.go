package main

import (
	"github.com/dknathalage/modules/invoice/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", web.Index)
	e.HideBanner = true

	e.Logger.Fatal(e.Start(":80"))
}
