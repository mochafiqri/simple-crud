package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mochafiqri/simple-crud/controllers"
	"github.com/mochafiqri/simple-crud/infrastructures"
	"github.com/mochafiqri/simple-crud/repository"
	"github.com/mochafiqri/simple-crud/usecases"
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

	var repoContent = repository.NewContentRepo(db, rds)
	var ucContent = usecases.NewContentUseCase(repoContent)
	var contentApi = controllers.NewHandler(ucContent)
	contentApi.Routes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
