package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	var e = echo.New()
	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "hello word",
		})
	})
	e.Logger.Fatal(e.Start(":8080"))
}
