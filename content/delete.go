package content

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Delete(e echo.Context) error {
	var id = e.Param("id")

	//find content
	if _, ok := Contents[id]; !ok {
		return e.JSON(http.StatusBadRequest, Resp{
			Message: "content not found",
			Data:    nil,
		})
	}

	delete(Contents, id)

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	})
}
