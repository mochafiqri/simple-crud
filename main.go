package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mochafiqri/simple-crud/content"
	"net/http"
)

func main() {
	var e = echo.New()
	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "hello word",
		})
	})

	e.GET("/content", content.ReadAll)
	e.GET("/content/:id", content.ReadById)
	e.POST("/content", content.Create)
	e.PUT("/content/:id", content.Update)
	e.DELETE("/content/:id", content.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
