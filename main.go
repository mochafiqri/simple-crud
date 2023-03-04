package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mochafiqri/simple-crud/controllers"
	"net/http"
)

func main() {
	var e = echo.New()
	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "hello word",
		})
	})

	e.GET("/contents", controllers.ReadAll)
	e.GET("/contents/:id", controllers.ReadById)
	e.POST("/contents", controllers.Create)
	e.PUT("/contents/:id", controllers.Update)
	e.DELETE("/contents/:id", controllers.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
