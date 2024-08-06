package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	/*
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
	*/
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "public",
		Browse:     true,
		IgnoreBase: true,
	}))

	e.File("/", "public/index.html")
	e.Logger.Fatal(e.Start(":1323"))
}
