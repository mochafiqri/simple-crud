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

	e.GET("/controllers", controllers.ReadAll)
	e.GET("/controllers/:id", controllers.ReadById)
	e.POST("/controllers", controllers.Create)
	e.PUT("/controllers/:id", controllers.Update)
	e.DELETE("/controllers/:id", controllers.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
