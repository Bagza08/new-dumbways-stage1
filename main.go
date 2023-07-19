package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hallo world")
	})

	e.GET("/json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"name":    "Bagza",
			"address": "Bandung",
		})
	})

	e.Logger.Fatal(e.Start("localhost:5000"))
}
