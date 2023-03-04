package content

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ReadAll(e echo.Context) error {
	var (
		contents = make([]Content, 0)
	)

	if len(Contents) == 0 {
		return e.JSON(http.StatusNotFound, Resp{
			Message: "data not found",
			Data:    nil,
		})
	}

	for _, v := range Contents {
		contents = append(contents, v)
	}

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    contents,
	})
}

func ReadById(e echo.Context) error {
	var id = e.Param("id")

	if _, ok := Contents[id]; !ok { // if content not found
		return e.JSON(http.StatusNotFound, Resp{
			Message: "data not found",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, Resp{
		Message: http.StatusText(http.StatusOK),
		Data:    Contents[id],
	})

}
