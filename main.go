package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mochafiqri/simple-crud/controllers"
	"github.com/mochafiqri/simple-crud/infrastructures"
	"net/http"
)

func main() {
	var e = echo.New()
	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "hello word",
		})
	})

	db, err := infrastructures.InitMysql()
	if err != nil {
		panic(err)
	}

	rds, err := infrastructures.InitRedis()
	if err != nil {
		panic(err)
	}

	var h = controllers.Handler{
		Db:  db,
		Rds: rds,
	}

	e.GET("/contents", h.ReadAll)
	e.GET("/contents/redis/flush", func(c echo.Context) error {
		err = rds.FlushAll(context.Background()).Err()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
	})
	e.GET("/contents/:id", h.ReadById)
	e.POST("/contents", h.Create)
	e.PUT("/contents/:id", h.Update)
	e.DELETE("/contents/:id", h.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
